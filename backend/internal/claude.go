package internal

import (
	"errors"

	"github.com/h00s/tinylink/app/models"
	"github.com/madebywelch/anthropic-go/v2/pkg/anthropic"
)

type Claude struct {
	client *anthropic.Client
}

func NewClaude(apiKey string) (*Claude, error) {
	if client, err := anthropic.NewClient(apiKey); err == nil {
		return &Claude{
			client: client,
		}, nil
	} else {
		return nil, err
	}
}

func (c *Claude) CheckGrammar(content models.Content) (models.Content, error) {
	var prompt string

	switch content.Lang {
	case "hr":
		prompt = "Provjeri gramatiku u sljedećem tekstu i odgovori samo sa ispravljenim tekstom bez ikakvog uvoda ili odgovori izvornim tekstom ukoliko nije potreban ispravak: "
	case "en":
		prompt = "Check the grammar in the following text and respond only with the corrected text without any introduction or respond with the original text if no correction is needed: "
	default:
		return content, errors.New("unsupported language")
	}

	prompt = prompt + content.Text
	request := anthropic.NewMessageRequest(
		[]anthropic.MessagePartRequest{{Role: "user", Content: []anthropic.ContentBlock{anthropic.NewTextContentBlock(prompt)}}},
		anthropic.WithModel[anthropic.MessageRequest](anthropic.Claude3Haiku),
		anthropic.WithMaxTokens[anthropic.MessageRequest](1000),
	)

	response, err := c.client.Message(request)
	if err != nil {
		return content, err
	}

	content.Text = response.Content[0].Text
	return content, nil
}
