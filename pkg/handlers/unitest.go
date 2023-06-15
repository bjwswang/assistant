package handlers

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/bjwswang/assistant/pkg/assistant"
	"github.com/gofiber/fiber/v2"
)

type UnitTestHandler struct {
	*assistant.Assistant
}

// NewUnitTestHandler creates a new instance of the UnitTestHandler
func NewUnitTestHandler(ai *assistant.Assistant) *UnitTestHandler {
	return &UnitTestHandler{
		Assistant: ai,
	}
}

// UnitTestRequest represents the request body for generating unit tests for a given code
type UnitTestRequest struct {
	Code string `json:"code"`
}

// enerateUnitTests generates unit tests
func (handler *UnitTestHandler) GenerateUnitTests(ctx *fiber.Ctx) error {
	var unitTestRequest UnitTestRequest
	if err := ctx.BodyParser(&unitTestRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validateUnitTestRequest(&unitTestRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	decodeCode, err := base64.StdEncoding.DecodeString(unitTestRequest.Code)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	resp, err := handler.ForUnitTests().Call(context.Background(), map[string]any{
		"code": string(decodeCode),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(resp)
}

func validateUnitTestRequest(unitTestRequest *UnitTestRequest) error {
	if unitTestRequest.Code == "" {
		return errors.New("code is required")
	}
	return nil
}
