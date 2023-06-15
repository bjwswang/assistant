package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/bjwswang/assistant/pkg/assistant"
	"github.com/bjwswang/assistant/pkg/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"k8s.io/klog/v2"
)

var (
	cfgFile = flag.String("config", "assistant.json", "config file path")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		klog.Error(err)
	}
}

// run starts the server and initializes the contract client
func run() error {
	klog.Infoln("Creating http server")

	// read from config file and load into fiber.Config
	cfgData, err := os.ReadFile(*cfgFile)
	if err != nil {
		return err
	}
	var config = Config{
		Addr:      ":9999",
		Assistant: assistant.DefaultConfig(),
	}
	err = json.Unmarshal(cfgData, &config)
	if err != nil {
		return err
	}

	// initialize assistant
	aiAssistant := assistant.New(&config.Assistant)

	// create a new fiber app
	app := fiber.New(config.Fiber)
	// add CORS middleware
	app.Use(cors.New(cors.ConfigDefault))
	// add logger middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// add routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to AI Assistant 👋! \n")
	})
	app.Post("/chat", handlers.NewChatHandler(aiAssistant).Chat)
	app.Post("/ut", handlers.NewUnitTestHandler(aiAssistant).GenerateUnitTests)

	klog.Infoln("Starting assistant server")
	if err := app.Listen(config.Addr); err != nil {
		return err
	}

	return nil
}
