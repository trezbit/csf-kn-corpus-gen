package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/trezbit/csf-kn-corpus-generator/chatgpti"
	"github.com/trezbit/csf-kn-corpus-generator/model"
	"github.com/trezbit/csf-kn-corpus-generator/utils"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

// runWorkerGroupSession is a function that runs a worker group session to build and process a batch of corpus generation queries
// issuing them to ChatGPT asychronously. The function takes in the query type (assessment, core, mapping), input and output paths, and the ChatGPT client.
func RunWorkerGroupSession(qtype string, inpath string, outpath string, c *openai.Client) {
	// build queries
	qconverse, err := model.BuildCorpusQueries(qtype, inpath, outpath)
	utils.CheckError(err)

	var wg sync.WaitGroup
	_batch_size := viper.GetInt("API_BATCH_SIZE")
	_wait_seconds := viper.GetInt("API_WAIT_SECONDS")

	for i, qconv := range qconverse {

		if i > 1 && i%_batch_size == 0 {
			time.Sleep(time.Duration(_wait_seconds) * time.Second)
		}
		wg.Add(1)
		// Call the function
		fmt.Println(i, qconv.QContent)

		go func(query string, outpath string, c *openai.Client) {
			defer wg.Done()
			err := chatgpti.GetGenResponse(query, outpath, c)
			utils.CheckError(err)
		}(qconv.QContent, qconv.RPath, c)

	}

	wg.Wait()
}

// Command line utility to generate a corpus of questions, and responses from ChatGPT for a given set of CSF controls
// and known mappings to other standards. The utility can be used to generate a corpus for:
// 1 - Training domain specific models and KBs that would allow mapping between different standards
// 2 - To generate a questionnaire corpus for assessing compliance with a standard.
// 3 - To build KNNs inferring mappings and trace continuity between different versions of CSF.
func main() {

	// load in config
	viper.SetConfigFile("config/.env")
	viper.ReadInConfig()

	// Get ChatGPT client
	apikey := os.Getenv("OPEN_API_SECRET")
	org := os.Getenv("OPEN_API_ORGID")
	c := chatgpti.NewGptClient(apikey, org)

	// Check args for the type of processing
	var qtype string
	flag.StringVar(&qtype, "qtype", "csf", "Query type [csf|assess|map]")
	flag.Parse()
	var inpath string
	var outpath string

	if qtype == "assess" || qtype == "assessq" {
		fmt.Println(viper.Get("INPATH_QUERY_ASSESS"))
		inpath = viper.Get("INPATH_QUERY_ASSESS").(string)
		outpath = viper.Get("OUTPATH_ASSESS").(string)
	} else if qtype == "csf" {
		fmt.Println(viper.Get("INPATH_QUERY_CSF"))
		inpath = viper.Get("INPATH_QUERY_CSF").(string)
		outpath = viper.Get("OUTPATH_CSF").(string)
	} else if qtype == "csfkey" {
		fmt.Println(viper.Get("INPATH_QUERY_CSF"))
		inpath = viper.Get("INPATH_QUERY_CSF").(string)
		outpath = viper.Get("OUTPATH_CSF").(string)
	} else if qtype == "map" {
		fmt.Println(viper.Get("INPATH_QUERY_MAP"))
		inpath = viper.Get("INPATH_QUERY_MAP").(string)
		outpath = viper.Get("OUTPATH_MAP").(string)
	} else {
		fmt.Println("Unrecognized/Unimplemented Query Type:", qtype)
		os.Exit(1)
	}
	// dispatch session
	RunWorkerGroupSession(qtype, inpath, outpath, c)

}
