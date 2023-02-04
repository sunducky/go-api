package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sunducky/go-api-template/config"
	"github.com/sunducky/go-api-template/pkg/shutdown"
)

func main() {
	// Setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// Load config
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// Run the server
	cleanup, err := run(env)

	// Run cleanup after server is terminated
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// Ensure server is shutdown gracefully and app runs
	shutdown.Gracefully()

}

func run(env config.EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	// start server
	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	// return a dunction to close server and database
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) (*fiber.App, func(), error) {
	// we initialize our storage here

	// create fiber app
	app := fiber.New()

	// We add global middlewares here
	app.Use(cors.New())
	app.Use(logger.New())

	// Add helth check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("🦆 Papa_Quacky is healthy 🦆")
	})

	// Add swagger docs endpoint here, when i finally decide to do
	// TODO

	// Register domains here

	return app, func() {
		// closeing of db
	}, nil

}
