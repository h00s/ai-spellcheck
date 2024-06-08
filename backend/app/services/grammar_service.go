package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/internal"
)

type GrammarService struct {
	raptor.Service

	claude *internal.Claude
}

func NewGrammarService() *GrammarService {
	ss := &GrammarService{}

	ss.OnInit(func() {
		var err error
		if ss.claude, err = internal.NewClaude(ss.Config.AppConfig["anthropic_key"].(string)); err != nil {
			ss.Log.Error("Error creating Claude", "error", err.Error())
		}
	})

	return ss
}
