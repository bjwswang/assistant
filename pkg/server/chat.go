package server

import (
	"context"
	"errors"

	"github.com/bjwswang/assistant/pkg/assistant"
	"github.com/gofiber/fiber/v2"
	"github.com/tmc/langchaingo/schema"
)

type ChatHandler struct {
	*assistant.Assistant
}

func NewChatHandler(ai *assistant.Assistant) *ChatHandler {
	return &ChatHandler{
		ai,
	}
}

type ChatRequest struct {
	// Question to ask assistant
	Question string `json:"question"`
	// Args to pass to assistant
	Args []string `json:"args,omitempty"`
}

type ChatResponse struct {
	Answer string `json:"answer"`
}

func (handler *ChatHandler) Chat(ctx *fiber.Ctx) error {
	// parse request body into struct
	var chatRequest ChatRequest
	if err := ctx.BodyParser(&chatRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate request
	if err := validateChatRequest(&chatRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// get answer from assistant
	completion, err := handler.ForChat().Chat(context.Background(), []schema.ChatMessage{
		schema.HumanChatMessage{
			Text: chatRequest.Question,
		},
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(ChatResponse{
		Answer: completion.Message.GetText(),
	})
}

// validateChatRequest validates the chat request
func validateChatRequest(chatRequest *ChatRequest) error {
	// validate question
	if chatRequest.Question == "" {
		return errors.New("question is required")
	}
	// validate args
	if len(chatRequest.Args) > 0 {
		for _, arg := range chatRequest.Args {
			if arg == "" {
				return errors.New("args must not be empty")
			}
		}
	}
	return nil
}
