package chatgpti

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/trezbit/csf-kn-corpus-generator/utils"
)

func NewGptClient(apikey string, org string) *openai.Client {
	config := openai.DefaultConfig(apikey)
	config.OrgID = org
	c := openai.NewClientWithConfig(config)
	return c
}

func GetGenResponse(query string, outpath string, c *openai.Client) error {
	fmt.Println("Processing: ", query, outpath)

	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// To-DO: make the model selection configurable
			// e.g., Model: openai.GPT3Dot5Turbo,
			Model: openai.GPT4Turbo0125,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: query,
				},
			},
		},
	)

	utils.CheckError(err)

	f, err := os.Create(outpath)
	utils.CheckError(err)
	defer f.Close()
	n3, err := f.WriteString(resp.Choices[0].Message.Content)
	utils.CheckError(err)
	_ = n3
	f.Sync()

	return nil
}
