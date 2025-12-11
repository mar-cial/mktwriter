package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ollama/ollama/api"
)

func main() {
	company, err := loadCompany()
	if err != nil {
		log.Fatalln(err)
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalln(err)
	}

	aiService := NewAIService(client, "gpt-oss:120b-cloud", company)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	prompt := strings.Join(os.Args[1:], " ")
	if prompt == "" {
		log.Fatalln("usage: provide a basic idea of what you want the post to be about related to your company")
	}

	response, err := aiService.MarketingBot(ctx, prompt)
	if err != nil {
		log.Fatalln(err)
	}

	if err := StorePost(response); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n\nPost was generated!\n\n")
}
