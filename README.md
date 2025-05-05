# AI Chatbot Using DeepSeek API

This project implements an AI-powered chatbot using the DeepSeek API for natural language understanding. It uses Go and the `langchaingo` library to interact with the API, providing a simple command-line interface to chat with the AI.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Setup](#setup)
- [Running the Chatbot](#running-the-chatbot)
- [Environment Variables](#environment-variables)
- [Usage](#usage)
- [Contributing](#contributing)

## Prerequisites

- Go (1.18+)
- DeepSeek API Key (create an account at [DeepSeek](https://platform.deepseek.com/))
- `go get` dependencies for the `langchaingo` library
- An internet connection for API communication

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/deepseek-chatbot.git
    cd LLM-CLI
    ```

2. Initialize the Go project and install dependencies:

    ```bash
    go get github.com/tmc/langchaingo/llms
    go get github.com/joho/godotenv
    github.com/spf13/cobra@latest
    go install github.com/spf13/cobra-cli@latest
    ```

3. Download and install the required dependencies:

    ```bash
    go mod tidy
    ```

## Setup

1. Create a `.env` file in the project root directory to store your **DeepSeek API key**:

    ```txt
    DEEPSEEK_API_KEY=your_deepseek_api_key_here
    ```

2. Ensure your Go environment is set up correctly. You can verify this by running:

    ```bash
    go version
    ```

## Running the Chatbot

1. Once you’ve set up the project and installed dependencies, you can run the chatbot via the command line:

    ```bash
    go run main.go chat
    ```

2. When prompted, enter an initial system prompt for the AI to understand how it should behave. Example:

    ```txt
    You are an AI developer and you have to answer my queries
    ```

3. Start chatting with the AI! Type your questions and the chatbot will respond using DeepSeek's language model.

4. To exit the chat, type `quit` or `exit`.

## Environment Variables

- **DEEPSEEK_API_KEY**: Set your DeepSeek API key in the `.env` file.

    Example:

    ```txt
    DEEPSEEK_API_KEY=your_deepseek_api_key_here
    ```

## Usage

- **Initial Prompt**: When you run the chatbot, it will ask for an initial prompt. This prompt sets the AI's tone for the conversation.
- **Chat**: After providing the prompt, you can start typing your queries, and the AI will respond.
- **Exit**: Type `quit` or `exit` to end the conversation.
