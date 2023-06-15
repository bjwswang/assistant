package assistant

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/outputparser"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"k8s.io/klog"
)

var (
	ChainUnitTestInputKeys  = []string{"code"}
	ChainUnitTestOutputKeys = []string{"unit_tests"}
)

const _gernerateUnitTestsTempalte = `
Generate unit tests for the following code:

{{.code}}

Requirements:
1. The unit tests should have >80% code coverage.

Requirements for the response:
1. The reponse should be well formatted which is ready to run.
`

var _ chains.Chain = (*ChatUnitTest)(nil)

type ChatUnitTest struct {
	config ModelConfig

	prompt prompts.PromptTemplate
	llm    ComposedLLM

	memory schema.Memory

	outputParser schema.OutputParser[any]

	inputKeys  []string
	outputKeys []string
}

func NewChainUnitTest(config ModelConfig, llm ComposedLLM) chains.Chain {
	var chainUnitTest = &ChatUnitTest{
		config:       config,
		llm:          llm,
		memory:       memory.NewSimple(),
		outputParser: outputparser.NewSimple(),
		inputKeys:    ChainUnitTestInputKeys,
		outputKeys:   ChainUnitTestOutputKeys,
	}
	template := prompts.NewPromptTemplate(_gernerateUnitTestsTempalte, chainUnitTest.inputKeys)

	chainUnitTest.prompt = template

	return chainUnitTest
}

func (c *ChatUnitTest) Call(ctx context.Context, values map[string]any, options ...chains.ChainCallOption) (map[string]any, error) {
	promptValue, err := c.prompt.FormatPrompt(values)
	if err != nil {
		return nil, err
	}
	klog.Infof("Call ChatUnitTest with prompt: %s \n", promptValue.String())
	klog.Infof("Call ChatUnitTest with model %s \n", c.config.Model)

	switch c.config.Model {
	case Davinci:
		generations, err := c.llm.Generate(ctx, []string{promptValue.String()})
		if err != nil {
			return nil, err
		}
		finalOutput, err := c.outputParser.ParseWithPrompt(generations[0].Text, promptValue)
		if err != nil {
			return nil, err
		}
		return map[string]any{c.outputKeys[0]: finalOutput}, nil
	case GPT3:
		generations, err := c.llm.Chat(context.Background(), []schema.ChatMessage{
			schema.SystemChatMessage{
				Text: "Think of yourself as a senior software developer.",
			},
			schema.HumanChatMessage{
				Text: promptValue.String(),
			},
		}, llms.WithTemperature(c.config.Temperature), llms.WithMaxTokens(c.config.MaxTokens))
		if err != nil {
			return nil, err
		}
		return map[string]any{c.outputKeys[0]: generations.Message.GetText()}, nil
	default:
		return nil, fmt.Errorf("unsupported model: %s", c.config.Model)
	}
}

// implements chains.Chain for ChainUnitTest
func (c *ChatUnitTest) GetInputKeys() []string {
	return c.inputKeys
}

// implements chains.Chain for ChainUnitTest
func (c *ChatUnitTest) GetOutputKeys() []string {
	return c.outputKeys
}

// implements chains.Chain for ChainUnitTest
func (c *ChatUnitTest) GetMemory() schema.Memory {
	return c.memory
}
