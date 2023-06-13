package assistant

import (
	"github.com/tmc/langchaingo/llms"
	langchain "github.com/tmc/langchaingo/llms/openai"
)

const (
	defaultOpenAIModel = "gpt-3.5-turbo"
)

type Config struct {

	// APIKey to access assistant api
	APIKey string `json:"api_key" validate:"required"`
	// Model to use for assistant
	// Use gpt-3.5-turbo by default
	Model string `json:"model"`
}

// Assistant is a wrapper around the assistant api
type Assistant struct {
	llm *langchain.LLM
}

// New creates a new assistant
func New(cfg *Config) *Assistant {
	if cfg.Model == "" {
		cfg.Model = defaultOpenAIModel
	}
	assistant := &Assistant{}

	// initialize a openai client with tmc/langchaingo
	llm, err := langchain.New(
		langchain.WithToken(cfg.APIKey),
		langchain.WithModel(cfg.Model),
	)
	if err != nil {
		panic("failed to initialize assistant openai client: " + err.Error())
	}
	assistant.llm = llm

	return assistant
}

func (assistant *Assistant) ForChat() llms.ChatLLM {
	return assistant.llm
}
