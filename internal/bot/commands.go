package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// handleCommand processes incoming commands from the user
func handleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		// Send a welcome sticker and message with a keyboard
		sendSticker(bot, msg.Chat.ID, "CAACAgIAAxkBAUnfw2YJcAxJpi6T9NHd8LsJkYTq_eQGAAIlAANd6qsi6WHKxUajPyQ0BA")
		sendMessageWithKeyboard(bot, msg.Chat.ID, "Let's play Rock Paper Scissors! Choose your move:", createKeyboard())
	case "bye":
		// Send a goodbye sticker and message
		sendStickerAndMessage(bot,
			msg.Chat.ID, "CAACAgIAAxkBAUnfwGYJb_fw-cYOf7_g790oVUaEz_OTAAInAANd6qsiTtaS6Yvg0mU0BA",
			"Bye, thanks for playing. Press /start to wake me up!")
	default:
		// Inform the user about the invalid command
		sendMessage(bot, msg.Chat.ID, "Invalid command. Use /start to begin.")
	}
}
