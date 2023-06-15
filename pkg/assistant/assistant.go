package assistant

import (
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type Model string

const (
	GPT3    Model = "gpt-3.5-turbo"
	Davinci Model = "text-davinci-003"
)

func (m Model) String() string {
	return string(m)
}

var (
	defaultChatOpenAIModel = ModelConfig{
		Model:       GPT3,
		Temperature: 0.7,
		MaxTokens:   150,
	}
	defaultCompletionOpenAIModel = ModelConfig{
		Model:       GPT3,
		Temperature: 0.5,
		MaxTokens:   150,
	}
)

type Config struct {
	// APIKey to access assistant api
	APIKey   string      `json:"api_key"`
	Chat     ModelConfig `json:"chat"`
	UnitTest ModelConfig `json:"unit_test"`
}

// Config for assistant
type ModelConfig struct {
	// Model to use for assistant
	// Use gpt-3.5-turbo by default
	Model `json:"model"`
	// Temperature for assistant by default
	Temperature float64 `json:"temperature"`
	// MaxTokens for assistant by default
	MaxTokens int `json:"max_tokens"`
}

func DefaultConfig() Config {
	return Config{
		Chat:     defaultChatOpenAIModel,
		UnitTest: defaultCompletionOpenAIModel,
	}
}

type ComposedLLM interface {
	llms.LLM
	llms.ChatLLM
}

// Assistant is a wrapper around the assistant api
type Assistant struct {
	// config for assistant
	cfg *Config
	// llm for chat
	chatllm *openai.LLM
	// llm for unit test
	utchain chains.Chain
}

// New creates a new assistant
func New(cfg *Config) *Assistant {
	if cfg.APIKey == "" {
		panic("api_key is required")
	}

	assistant := &Assistant{
		cfg: cfg,
	}

	chatllm, err := openai.New(
		openai.WithToken(cfg.APIKey),
		openai.WithModel(cfg.Chat.Model.String()),
	)
	if err != nil {
		panic("failed to initialize assistant chat client: " + err.Error())
	}
	assistant.chatllm = chatllm

	utllm, err := openai.New(
		openai.WithToken(cfg.APIKey),
		openai.WithModel(cfg.UnitTest.Model.String()),
	)
	if err != nil {
		panic("failed to initialize assistant completion llm client: " + err.Error())
	}
	assistant.utchain = NewChainUnitTest(cfg.UnitTest, utllm)

	return assistant
}

// Config returns the config for assistant
func (assistant *Assistant) ChatConfig() ModelConfig {
	return assistant.cfg.Chat
}
func (assistant *Assistant) CompletionConfig() ModelConfig {
	return assistant.cfg.UnitTest
}

func (assistant *Assistant) ForChat() ComposedLLM {
	return assistant.chatllm
}

func (assistant *Assistant) ForUnitTests() chains.Chain {
	return assistant.utchain
}
