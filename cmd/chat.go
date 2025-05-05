/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "LLM Chatbot",
	Long:  `LLM Chatbot`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load()
		reader := bufio.NewReader(os.Stdin)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigChan
			fmt.Println("\nInterrupt signal received. Exiting...")
			os.Exit(0)
		}()

		llm, err := openai.New(
			openai.WithBaseURL("https://api.deepseek.com/v1"),
			openai.WithToken(os.Getenv("DEEPSEEK_API_KEY")),
			openai.WithModel("deepseek-chat"),
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Enter intial prompt for LLM: ")
		intialPrompt, _ := reader.ReadString('\n')
		intialPrompt = strings.TrimSpace(intialPrompt)
		content := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, intialPrompt),
		}
		fmt.Print("Prompt received. Entering chat mode...")

		ctx := context.Background()

		for {
			fmt.Println("> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			switch input {
			case "quit", "exit", "Quit", "Exit":
				fmt.Print("Exiting...")
				os.Exit(0)
			default:
				response := ""
				content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, input))
				_, err := llm.GenerateContent(ctx, content,
					llms.WithMaxTokens(1024),
					llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
						fmt.Print(string(chunk))
						response = response + string(chunk)
						return nil
					}),
				)
				if err != nil {
					fmt.Println("\n[Error while generating response]:", err)
					continue
				}
				content = append(content, llms.TextParts(llms.ChatMessageTypeSystem, response))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
