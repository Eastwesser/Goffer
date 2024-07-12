package bot

import (
	"log"
	"math/rand"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

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

func sendMessageWithKeyboard(bot *tgbotapi.BotAPI, chatID int64, text string, keyboard tgbotapi.ReplyKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func sendStickerAndMessage(bot *tgbotapi.BotAPI, chatID int64, stickerID, messageText string) {
	sendSticker(bot, chatID, stickerID)
	sendMessage(bot, chatID, messageText)
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func sendSticker(bot *tgbotapi.BotAPI, chatID int64, stickerID string) {
	msg := tgbotapi.NewStickerShare(chatID, stickerID)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending sticker:", err)
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
