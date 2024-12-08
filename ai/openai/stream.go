// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// By accessing or using this software, you agree to be bound by the terms
// of the License Agreement, which you can find at LICENSE files.

//go:build !json_sonic
// +build !json_sonic

package openai

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/valyala/fasthttp"
)

// Client represents a client to interact with the OpenAI API.
type Client struct {
	APIKey string
}

// NewClient initializes a new OpenAIClient using the API key from environment variables.
// Panics if the API key is not set.
func NewClient() *Client {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("API key not set")
	}
	return &Client{APIKey: apiKey}
}

// ChatCompletionsStreamResponse represents the structure of a streaming response from the OpenAI API.
type ChatCompletionsStreamResponse struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}

// StreamChatCompletion handles streaming chat completions from the OpenAI API.
// It sends a POST request to the OpenAI API and streams the response back to the client.
//
// Note: This is just an example. If you have your own AI, it's much easier to create own SDK written in Go for this SSE.
func (ai *Client) StreamChatCompletion(c *fiber.Ctx) error {
	// List of random user messages
	//
	// Note: This is just an example to simulate user requests. In production, these messages would come from the user's request body.
	userMessages := []string{
		"Hello!",
		"How's the weather today?",
		"Tell me a joke.",
		"What's the latest news?",
		"Can you help me with a math problem?",
	}

	// Select a random message using [crypto/rand]
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(userMessages))))
	if err != nil {
		log.Printf("Failed to generate random index: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate random index")
	}
	randomMessage := userMessages[index.Int64()]

	reqBody := map[string]any{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": randomMessage},
		},
		"stream": true,
	}

	reqBytes, err := c.App().Config().JSONEncoder(reqBody)
	if err != nil {
		log.Printf("Failed to encode request body: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to encode request body")
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI("https://api.openai.com/v1/chat/completions")
	req.Header.SetMethod("POST")
	// Note: Using the content-type text/event-stream with OpenAI may not work anymore.
	// If it still works, it can be more challenging to manage.
	// JSON is preferred because it's easier to stream back to the client and can be improved due to its implementation in Go.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ai.APIKey))
	req.SetBody(reqBytes)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	clientHTTP := &fasthttp.Client{}
	if err = clientHTTP.Do(req, resp); err != nil {
		log.Printf("Request failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Request failed")
	}

	if resp.StatusCode() != fiber.StatusOK {
		return c.Status(resp.StatusCode()).SendString(string(resp.Body()))
	}

	// Note: This is just an example and needs improvement for production use.
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	// Note: Handling the OpenAI stream response with resp.BodyStream in FastHTTP is not possible,
	// as it will always return nil due to differences in how streaming works.
	body := resp.Body()
	reader := bufio.NewReader(strings.NewReader(string(body)))

	c.Context().Response.SetBodyStreamWriter(func(w *bufio.Writer) {
		var fullMessage strings.Builder

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Error reading line: %v", err)
				break
			}

			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			if line == "data: [DONE]" {
				break
			}

			if strings.HasPrefix(line, "data: ") {
				line = strings.TrimPrefix(line, "data: ")

				var streamResponse ChatCompletionsStreamResponse
				if err := c.App().Config().JSONDecoder([]byte(line), &streamResponse); err != nil {
					log.Printf("Error parsing JSON: %v", err)
					continue
				}

				for _, choice := range streamResponse.Choices {
					content := choice.Delta.Content
					if content != "" {
						fullMessage.WriteString(content)
						w.WriteString(fmt.Sprintf("data: %s\n\n", content))
						w.Flush()
					}
				}
			}
		}

		log.Printf("Full Message: %s", fullMessage.String())
	})

	return nil
}
