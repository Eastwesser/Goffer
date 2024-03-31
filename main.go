package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
)

var bot *tgbotapi.BotAPI

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the bot token from the environment variables
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN not found in environment variables")
	}

	// Create a new bot instance
	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Start the main loop to handle updates
	startMainLoop()
}

func startMainLoop() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			handleCommand(update.Message)
		} else {
			handleMessage(update.Message)
		}
	}
}

func handleCommand(msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		sendMessage(msg.Chat.ID, "Welcome to Rock Paper Scissors! Use /play to start the game.")
	case "play":
		startGame(msg.Chat.ID)
	default:
		sendMessage(msg.Chat.ID, "Invalid command. Use /start to begin.")
	}
}

func handleMessage(msg *tgbotapi.Message) {
	if msg.Text == "rock" || msg.Text == "paper" || msg.Text == "scissors" {
		handleGameAction(msg, msg.Text) // Pass user's choice to handleGameAction
	} else {
		sendMessage(msg.Chat.ID, "I'm sorry, I didn't understand that.")
	}
}

func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func startGame(chatID int64) {
	sendMessage(chatID, "Let's play Rock Paper Scissors! Choose your move: rock, paper, or scissors.")
}

func handleGameAction(msg *tgbotapi.Message, userChoice string) { // Add userChoice parameter
	botChoice := generateBotChoice()
	result := compareChoices(userChoice, botChoice)

	switch result {
	case "win":
		sendMessage(msg.Chat.ID, fmt.Sprintf("You chose %s, I chose %s. You win!", userChoice, botChoice))
	case "lose":
		sendMessage(msg.Chat.ID, fmt.Sprintf("You chose %s, I chose %s. You lose!", userChoice, botChoice))
	default:
		sendMessage(msg.Chat.ID, fmt.Sprintf("You chose %s, I chose %s. It's a tie!", userChoice, botChoice))
	}
}

func generateBotChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.Intn(len(choices))]
}

func compareChoices(userChoice, botChoice string) string {
	if userChoice == botChoice {
		return "tie"
	}
	switch userChoice {
	case "rock":
		if botChoice == "scissors" {
			return "win"
		}
	case "paper":
		if botChoice == "rock" {
			return "win"
		}
	case "scissors":
		if botChoice == "paper" {
			return "win"
		}
	}
	return "lose"
}
