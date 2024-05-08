package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/trezbit/csf-kn-corpus-generator/chatgpti"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func testinit(outpath string) *openai.Client {
	// Get ChatGPT client
	apikey := os.Getenv("OPEN_API_SECRET")
	org := os.Getenv("OPEN_API_ORGID")
	c := chatgpti.NewGptClient(apikey, org)

	os.RemoveAll(outpath)
	os.Mkdir(outpath, 0755)

	return c
}

func testcleanup(outpath string) {
	os.RemoveAll(outpath)

}

func TestAssessCorpusGen(test *testing.T) {
	fmt.Println("Running: TestAssessCorpusGen")
	// load in config
	viper.SetConfigFile("tests/.env")
	viper.ReadInConfig()
	// Check args for the type of processing
	var inpath string = viper.Get("INPATH_QUERY_ASSESS").(string)
	var outpath string = viper.Get("OUTPATH_ASSESS").(string)
	c := testinit(outpath)

	// dispatch session
	RunWorkerGroupSession("assess", inpath, outpath, c)

	// get all the files with extension .txt in the outpath
	files, err := os.ReadDir(outpath)
	if err != nil || len(files) < 2 {
		test.Error(err)
	}
	testcleanup(outpath)
}

func TestCSFCorpusGen(test *testing.T) {
	fmt.Println("Running: TestCSFCorpusGen")
	// load in config
	viper.SetConfigFile("tests/.env")
	viper.ReadInConfig()
	// Check args for the type of processing
	var inpath string = viper.Get("INPATH_QUERY_CSF").(string)
	var outpath string = viper.Get("OUTPATH_CSF").(string)
	c := testinit(outpath)

	// dispatch session
	RunWorkerGroupSession("csf", inpath, outpath, c)

	// get all the files with extension .txt in the outpath
	files, err := os.ReadDir(outpath)
	if err != nil || len(files) < 2 {
		test.Error(err)
	}
	testcleanup(outpath)
}

func TestMapCorpusGen(test *testing.T) {
	// load in config
	fmt.Println("Running: TestMapCorpusGen")

	viper.SetConfigFile("tests/.env")
	viper.ReadInConfig()
	// Check args for the type of processing
	var inpath string = viper.Get("INPATH_QUERY_MAP").(string)
	var outpath string = viper.Get("OUTPATH_MAP").(string)
	c := testinit(outpath)

	// dispatch session
	RunWorkerGroupSession("map", inpath, outpath, c)

	// get all the files with extension .txt in the outpath
	files, err := os.ReadDir(outpath)
	if err != nil || len(files) < 2 {
		test.Error(err)
	}
	testcleanup(outpath)
}
