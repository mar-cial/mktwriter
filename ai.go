package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ollama/ollama/api"
)

type AIService struct {
	Client  *api.Client
	Model   string
	Company *Company
}

func NewAIService(client *api.Client, model string, company *Company) *AIService {
	return &AIService{Client: client, Model: model, Company: company}

}

func (ai *AIService) MarketingBot(ctx context.Context, prompt string) (string, error) {
	system := "You are a marketing bot that writes short but useful posts (no more than 3 sentences) for facebook for this company: " + ai.Company.About + ""
	genReq := &api.GenerateRequest{
		Model:  ai.Model,
		Prompt: prompt,
		Suffix: "---",
		System: system,
	}

	var response strings.Builder
	genFn := func(res api.GenerateResponse) error {
		_, err := response.WriteString(res.Response)
		fmt.Print(res.Response)
		return err
	}

	if err := ai.Client.Generate(ctx, genReq, genFn); err != nil {
		return "", nil
	}

	return response.String(), nil
}
