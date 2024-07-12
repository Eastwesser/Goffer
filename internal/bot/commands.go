package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		sendSticker(bot, msg.Chat.ID, "CAACAgIAAxkBAUnfw2YJcAxJpi6T9NHd8LsJkYTq_eQGAAIlAANd6qsi6WHKxUajPyQ0BA")
		sendMessageWithKeyboard(bot, msg.Chat.ID, "Let's play Rock Paper Scissors! Choose your move:", createKeyboard())
	case "bye":
		sendStickerAndMessage(bot,
			msg.Chat.ID, "CAACAgIAAxkBAUnfwGYJb_fw-cYOf7_g790oVUaEz_OTAAInAANd6qsiTtaS6Yvg0mU0BA",
			"Bye, thanks for playing. Press /start to wake me up!")
	default:
		sendMessage(bot, msg.Chat.ID, "Invalid command. Use /start to begin.")
	}
}
