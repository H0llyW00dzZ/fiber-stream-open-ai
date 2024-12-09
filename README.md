# Fiber Stream Example

This repository demonstrates a simple implementation of a streaming chat application using the Fiber web framework, built on top of FastHTTP, and the OpenAI API.

## Overview

The application serves a web page where users can initiate a chat with an AI assistant. It uses Fiber to handle HTTP requests and streams responses from the OpenAI API to the client.

> [!NOTE]
> This example requires further improvements, such as `reducing latency`, if used in `production`

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

## Compatibility

- [**Kubernetes**](https://kubernetes.io/): Supports Horizontal Pod Autoscaler (HPA) when using external AI services (`e.g., the current example implementation that interacts with OpenAI`). For an in-house AI solution (`e.g., own AI`), consider using Vertical Pod Autoscaler (VPA) for resource management.

> [!NOTE]
> It also depends on the ingress configuration. If you have extensive Kubernetes knowledge (`e.g., a Captain Kubernetes`), it can be managed easily

> [!TIP]
> For Horizontal Pod Autoscaler (HPA), it's recommended to use custom metrics based on connections instead of the default `CPU/Memory` metrics. This approach can help achieve scalability, such as reaching up to 1000 nodes to handle 1 billion connections.

## Security Considerations

The current example implementation of the `SSE` primarily involves client-side operations (`e.g., even if there are vulnerabilities, it only affect the browser of the who views it`), which inherently carry a `lower security risk`. However, it's crucial to address potential vulnerabilities to ensure the safety and integrity of the application:

1. **Cross-Site Scripting (XSS):**
   - Ensure that any data rendered in the HTML is properly sanitized to prevent XSS attacks.
   - Validate and sanitize server-side responses before sending them to the client.

2. **Content Security Policy (CSP):**
   - Implement a CSP in HTTP headers to restrict the sources from which resources can be loaded. This helps mitigate XSS and data injection attacks.

3. **Secure Fetch Requests:**
   - Use HTTPS for all API requests to ensure data encryption during transmission.
   - Validate server responses and handle errors gracefully in client-side JavaScript to avoid exposing sensitive information.

4. **Access Control:**
   - Protect sensitive endpoints with authentication and authorization mechanisms to ensure only authorized users can access them.

5. **Error Handling:**
   - Implement robust error handling to prevent the exposure of sensitive information through error messages.

6. **Environment Variables:**
   - Store sensitive data, such as API keys, in environment variables and never hard-code them into the application.

By addressing these considerations, the security of the application can be enhanced while maintaining its performance and functionality.

## Acknowledgments

- [Fiber](https://gofiber.io/) for the web framework.
- [FastHTTP](https://github.com/valyala/fasthttp) for high-performance HTTP handling.
- [OpenAI](https://openai.com/) for the API services.
- [Sonic](https://github.com/bytedance/sonic) for optimized JSON processing.
