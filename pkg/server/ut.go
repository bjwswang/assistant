package server

import "github.com/gofiber/fiber/v2"

type UnitTestHandler struct{}

// NewUnitTestHandler creates a new instance of the UnitTestHandler
func NewUnitTestHandler() *UnitTestHandler {
	return &UnitTestHandler{}
}

type Language string

const (
	Go Language = "go"
)

// UnitTestRequest represents the request body for generating unit tests for a given code
type UnitTestRequest struct {
	RawCode  []byte `json:"rawCode"`
	Language string `json:"language"`
}

// TODO: implement GenerateUnitTests generates unit tests
func (handler *UnitTestHandler) GenerateUnitTests(ctx *fiber.Ctx) error {
	// parse request body into struct
	var unitTestRequest UnitTestRequest
	if err := ctx.BodyParser(&unitTestRequest); err != nil {
		return err
	}
	// check code language

	return nil
}
