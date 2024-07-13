package bot

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartMainLoop initializes the bot and starts the main update loop
func StartMainLoop(botToken string) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

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
			handleCommand(bot, update.Message)
		} else {
			handleMessage(bot, update.Message)
		}
	}
}
