package services

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

const TEXT_MODEL string = "gemini-2.5-flash"

func AITextSend(prompt string) (result string) {
	err := godotenv.Load()
	if (err != nil) {
		log.Fatalln("error loading ur env, is it created?")
	}

	ctx := context.Background()
	// client automatically gets "GEMINI_API_KEY" from .env
	client, err := genai.NewClient(ctx, nil)
	if (err != nil) {
		log.Println("error creating ai client. maybe 'GEMINI_API_KEY' is not in ur env")
		return err.Error()
	}

	resp, err := client.Models.GenerateContent(
		ctx,
		TEXT_MODEL,
		genai.Text(prompt),
		nil,
	)
	if (err != nil) {
		log.Println("error getting ai response")
		return err.Error()
	}
	result = resp.Text()
	return result
}