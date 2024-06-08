package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/internal"
)

type GrammarService struct {
	raptor.Service

	claude *internal.Claude
}

func NewGrammarService() *GrammarService {
	gs := &GrammarService{}

	gs.OnInit(func() {
		var err error
		if gs.claude, err = internal.NewClaude(gs.Config.AppConfig["anthropic_key"].(string)); err != nil {
			gs.Log.Error("Error creating Claude", "error", err.Error())
		}
	})

	return gs
}

func (gs *GrammarService) Check(content models.Content) (models.Content, error) {
	validatedContent, error := gs.claude.CheckGrammar(content)
	if error != nil {
		gs.Log.Error("Error checking grammar", "error", error.Error())
		return content, raptor.NewErrorInternal(error.Error())
	}
	return validatedContent, nil
}
