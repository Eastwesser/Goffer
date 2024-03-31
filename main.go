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
		sendSticker(msg.Chat.ID, "CAACAgIAAxkBAUnfw2YJcAxJpi6T9NHd8LsJkYTq_eQGAAIlAANd6qsi6WHKxUajPyQ0BA")
		sendMessageWithKeyboard(msg.Chat.ID, "Let's play Rock Paper Scissors! Choose your move:", createKeyboard())
	case "bye":
		sendStickerAndMessage(
			msg.Chat.ID, "CAACAgIAAxkBAUnfwGYJb_fw-cYOf7_g790oVUaEz_OTAAInAANd6qsiTtaS6Yvg0mU0BA",
			"Bye, thanks for playing. Press /start to wake me up!")
	default:
		sendMessage(msg.Chat.ID, "Invalid command. Use /start to begin.")
	}
}

func createKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Rock"),
			tgbotapi.NewKeyboardButton("Paper"),
			tgbotapi.NewKeyboardButton("Scissors"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Finish"),
		),
	)
	keyboard.OneTimeKeyboard = false // Keeping the keyboard visible
	return keyboard
}

func sendMessageWithKeyboard(chatID int64, text string, keyboard tgbotapi.ReplyKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func sendStickerAndMessage(chatID int64, stickerID, messageText string) {
	sendSticker(chatID, stickerID)
	sendMessage(chatID, messageText)
}

func handleMessage(msg *tgbotapi.Message) {
	switch msg.Text {
	case "Rock", "Paper", "Scissors":
		handleGameAction(msg, msg.Text)
	case "Finish":
		sendStickerAndMessage(
			msg.Chat.ID, "CAACAgIAAxkBAUnfwGYJb_fw-cYOf7_g790oVUaEz_OTAAInAANd6qsiTtaS6Yvg0mU0BA",
			"Bye, thanks for playing. Press /start to wake me up!")
	default:
		sendMessage(msg.Chat.ID, "I'm sorry, I didn't understand that. Type /start to wake me up!")
	}
}

func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func sendSticker(chatID int64, stickerID string) {
	msg := tgbotapi.NewStickerShare(chatID, stickerID)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending sticker:", err)
	}
}

func handleGameAction(msg *tgbotapi.Message, userChoice string) {
	botChoice := generateBotChoice()
	result := compareChoices(userChoice, botChoice)

	switch result {
	case "win":
		sendSticker(msg.Chat.ID, "CAACAgIAAxkBAUnfzWYJcBvWulVycGgv1TbxEopajrE3AAIXAANd6qsicsnBjr5cTb00BA")
		sendMessage(msg.Chat.ID, fmt.Sprintf("You chose %s, I chose %s. You win!", userChoice, botChoice))
	case "lose":
		sendSticker(msg.Chat.ID, "CAACAgIAAxkBAUnf0GYJcB2ERiExCjqYxebi4kR-1d2lAAIJAANd6qsi7-7sDc8Whpc0BA")
		sendMessage(msg.Chat.ID, fmt.Sprintf("You chose %s, I chose %s. You lose!", userChoice, botChoice))
	default:
		sendSticker(msg.Chat.ID, "CAACAgIAAxkBAUnf02YJcCup7gIIO5DMBND1PFZ3seDUAAIbAANd6qsinB_Cwhpp6Uo0BA")
		sendMessage(msg.Chat.ID, fmt.Sprintf("You chose %s, I chose %s. It's a tie!", userChoice, botChoice))
	}
}

func generateBotChoice() string {
	choices := []string{"Rock", "Paper", "Scissors"}
	return choices[rand.Intn(len(choices))]
}

func compareChoices(userChoice, botChoice string) string {
	if userChoice == botChoice {
		return "tie"
	}
	switch userChoice {
	case "Rock":
		if botChoice == "Scissors" {
			return "win"
		}
	case "Paper":
		if botChoice == "Rock" {
			return "win"
		}
	case "Scissors":
		if botChoice == "Paper" {
			return "win"
		}
	}
	return "lose"
}
