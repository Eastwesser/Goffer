package main

import (
	"log"
	"os"

	"github.com/Eastwesser/Goffer/internal/bot"
	"github.com/Eastwesser/Goffer/internal/redis"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redis.InitRedisClient()

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN not found in environment variables")
	}

	bot.StartMainLoop(botToken)
}
