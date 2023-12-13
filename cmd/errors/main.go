package main

import (
	"context"
	"fmt"
	"log"
	"os"

	gpt3 "github.com/chrisbward/go-gpt3orLLM"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEYS")
	if apiKey == "" {
		log.Fatalln("Missing OPENAI API KEYS")
	}

	ctx := context.Background()
	client := gpt3.NewClient(nil, []string{apiKey})

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt: []string{
			"1\n2\n3\n4",
		},
		MaxTokens: gpt3.IntPtr(0),
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", resp)
}
