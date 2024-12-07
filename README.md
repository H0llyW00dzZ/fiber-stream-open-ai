# Fiber Stream Example

This repository demonstrates a simple implementation of a streaming chat application using the Fiber web framework and the OpenAI API.

## Overview

The application serves a web page where users can initiate a chat with an AI assistant. It uses Fiber to handle HTTP requests and streams responses from the OpenAI API to the client.

## Features

- **HTML Template Rendering**: Uses Fiber's HTML template engine to render the interface.
- **Streaming API Integration**: Connects to the OpenAI API to stream chat responses.

## Project Structure

- **cmd/server/run.go**: Entry point of the application, sets up routes and handles server lifecycle.
- **frontend/views/**: Contains HTML templates for rendering the web interface.
- **ai/openai/**: Contains the client implementation for interacting with the OpenAI API.

## Key Functions

- **StreamChatCompletion**: Handles requests to the `/stream` endpoint, streaming chat responses from the OpenAI API to the client.
- **Graceful Shutdown**: Listens for OS signals to gracefully shut down the server, ensuring all ongoing requests are completed.

## Acknowledgments

- [Fiber](https://gofiber.io/) for the web framework.
- [OpenAI](https://openai.com/) for the API services.
