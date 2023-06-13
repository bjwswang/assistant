package main

import (
	"github.com/bjwswang/assistant/pkg/assistant"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	// Server address
	Addr string `json:"addr" validate:"required"`
	// Fiber config
	Fiber fiber.Config `json:"fiber"`
	// OpenAI config
	Assistant assistant.Config `json:"assistant" validate:"required,dive"`
}
