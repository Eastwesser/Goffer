package bot

import (
	"context"
	"fmt"
	"github.com/Eastwesser/Goffer/internal/redis"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// handleGameAction processes the user's game action and determines the result
func handleGameAction(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, userChoice string) {
	// Store the user's choice in Redis
	err := redis.Client.Set(context.Background(), fmt.Sprintf("user:%d:choice", msg.From.ID), userChoice, 0).Err()
	if err != nil {
		log.Println("Error storing user choice in Redis:", err)
	}

	// Generate the bot's choice and compare with the user's choice
	botChoice := generateBotChoice()
	result := compareChoices(userChoice, botChoice)

	// Send the appropriate response based on the game result and update highscore
	var errUpdate error
	switch result {
	case "win":
		sendStickerAndMessage(bot, msg.Chat.ID, "CAACAgIAAxkBAUnfzWYJcBvWulVycGgv1TbxEopajrE3AAIXAANd6qsicsnBjr5cTb00BA", fmt.Sprintf("You chose %s, I chose %s. You win!", userChoice, botChoice))
		errUpdate = redis.UpdateHighscore(int64(msg.From.ID), 1, 0, 0)
	case "lose":
		sendStickerAndMessage(bot, msg.Chat.ID, "CAACAgIAAxkBAUnf0GYJcB2ERiExCjqYxebi4kR-1d2lAAIJAANd6qsi7-7sDc8Whpc0BA", fmt.Sprintf("You chose %s, I chose %s. You lose!", userChoice, botChoice))
		errUpdate = redis.UpdateHighscore(int64(msg.From.ID), 0, 1, 0)
	default:
		sendStickerAndMessage(bot, msg.Chat.ID, "CAACAgIAAxkBAUnf02YJcCup7gIIO5DMBND1PFZ3seDUAAIbAANd6qsinB_Cwhpp6Uo0BA", fmt.Sprintf("You chose %s, I chose %s. It's a tie!", userChoice, botChoice))
		errUpdate = redis.UpdateHighscore(int64(msg.From.ID), 0, 0, 1)
	}

	if errUpdate != nil {
		log.Println("Error updating highscore:", errUpdate)
	}
}

// handleMessage processes incoming text messages from the user
func handleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	switch msg.Text {
	case "Rock", "Paper", "Scissors":
		// Handle the game action if the user sends a valid choice
		handleGameAction(bot, msg, msg.Text)
	case "Finish":
		// Send a goodbye sticker and message
		sendStickerAndMessage(bot,
			msg.Chat.ID, "CAACAgIAAxkBAUnfwGYJb_fw-cYOf7_g790oVUaEz_OTAAInAANd6qsiTtaS6Yvg0mU0BA",
			"Bye, thanks for playing. Press /start to wake me up!")
	default:
		// Inform the user about the invalid message
		sendMessage(bot, msg.Chat.ID, "I'm sorry, I didn't understand that. Type /start to wake me up!")
	}
}
