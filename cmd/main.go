package main

import (
	"log"
	"os"

	"github.com/Eastwesser/Goffer/internal/bot"
	"github.com/Eastwesser/Goffer/internal/redis"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the Redis client
	redis.InitRedisClient()

	// Get the bot token from the environment variables
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN not found in environment variables")
	}

	// Start the main update loop for the bot
	bot.StartMainLoop(botToken)
}
