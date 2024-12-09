// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// By accessing or using this software, you agree to be bound by the terms
// of the License Agreement, which you can find at LICENSE files.

//go:build json_sonic
// +build json_sonic

package main

import (
	"context"
	"fiber-stream/ai/openai"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize the HTML template engine
	engine := html.New("./frontend/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		// Note: If the response time for small text is around 2 seconds or less, it indicates stable.
		// To improve latency, it also depends on the configuration of Fiber and the FastHTTP client.
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	client := openai.NewClient()

	// Serve the HTML page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	// Stream endpoint
	app.Get("/stream", client.StreamChatCompletion)

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Start the server in a goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}()

	// Block until we receive a signal
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
