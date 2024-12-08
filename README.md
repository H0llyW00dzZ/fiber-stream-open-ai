# Fiber Stream Example

This repository demonstrates a simple implementation of a streaming chat application using the Fiber web framework, built on top of FastHTTP, and the OpenAI API.

## Overview

The application serves a web page where users can initiate a chat with an AI assistant. It uses Fiber to handle HTTP requests and streams responses from the OpenAI API to the client.

## Features

- **HTML Template Rendering**: Uses Fiber's HTML template engine to render the interface.
- **Streaming API Integration**: Connects to the OpenAI API to stream chat responses.
- **FastHTTP Performance**: Leverages FastHTTP for high-performance HTTP handling.
- **Sonic JSON Optimization**: Utilizes Sonic for efficient JSON encoding and decoding, reducing latency and improving performance.

## Project Structure

- **cmd/server/run.go**: Entry point of the application, sets up routes and handles server lifecycle.
- **frontend/views/**: Contains HTML templates for rendering the web interface.
- **ai/openai/**: Contains the client implementation for interacting with the OpenAI API, including optimized JSON handling with Sonic.

## Key Functions

- **StreamChatCompletion**: Handles requests to the `/stream` endpoint, streaming chat responses from the OpenAI API to the client.
- **Graceful Shutdown**: Listens for OS signals to gracefully shut down the server, ensuring all ongoing requests are completed.

## Sonic JSON Integration

- **Performance Enhancement**: Sonic is used for JSON encoding and decoding, which enhances performance by utilizing a pool of decoders and encoders.
- **Stream Handling**: The use of Sonic's `StreamDecoder` and `StreamEncoder` allows for efficient processing of streaming data, minimizing overhead.

## Acknowledgments

- [Fiber](https://gofiber.io/) for the web framework.
- [FastHTTP](https://github.com/valyala/fasthttp) for high-performance HTTP handling.
- [OpenAI](https://openai.com/) for the API services.
- [Sonic](https://github.com/bytedance/sonic) for optimized JSON processing.
