package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

// Redis client instance
var redisClient *redis.Client

// Highscore represents user highscore data
type Highscore struct {
	UserID int64
	Wins   int
	Losses int
	Draws  int
}

// Initialize Redis client
func initRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

// UpdateHighscore updates the user's highscore in Redis
func updateHighscore(userID int64, wins, losses, draws int) error {
	ctx := context.Background()
	key := fmt.Sprintf("highscores:%d", userID)
	err := redisClient.HSet(ctx, key, map[string]interface{}{
		"wins":   wins,
		"losses": losses,
		"draws":  draws,
	}).Err()
	if err != nil {
		log.Printf("Error updating highscore for user %d: %v", userID, err)
		return err
	}
	return nil
}

// GetHighscore retrieves the user's highscore from Redis
func getHighscore(userID int64) (Highscore, error) {
	ctx := context.Background()
	key := fmt.Sprintf("highscores:%d", userID)
	data, err := redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		log.Printf("Error retrieving highscore for user %d: %v", userID, err)
		return Highscore{}, err
	}
	hs := Highscore{
		UserID: userID,
	}
	if len(data) == 0 {
		return hs, nil
	}
	if val, ok := data["wins"]; ok {
		hs.Wins, _ = strconv.Atoi(val)
	}
	if val, ok := data["losses"]; ok {
		hs.Losses, _ = strconv.Atoi(val)
	}
	if val, ok := data["draws"]; ok {
		hs.Draws, _ = strconv.Atoi(val)
	}
	return hs, nil
}

// HandleGameAction handles the game action and updates highscore
func handleGameAction(msg *tgbotapi.Message, userChoice string) {
	err := redisClient.Set(context.Background(), fmt.Sprintf("user:%d:choice", msg.From.ID), userChoice, 0).Err()
	if err != nil {
		log.Println("Error storing user choice in Redis:", err)
	}

	botChoice := generateBotChoice()
	result := compareChoices(userChoice, botChoice)

	switch result {
	case "win":
		sendStickerAndMessage(msg.Chat.ID, "CAACAgIAAxkBAUnfzWYJcBvWulVycGgv1TbxEopajrE3AAIXAANd6qsicsnBjr5cTb00BA", fmt.Sprintf("You chose %s, I chose %s. You win!", userChoice, botChoice))
		updateHighscore(msg.From.ID, 1, 0, 0)
	case "lose":
		sendStickerAndMessage(msg.Chat.ID, "CAACAgIAAxkBAUnf0GYJcB2ERiExCjqYxebi4kR-1d2lAAIJAANd6qsi7-7sDc8Whpc0BA", fmt.Sprintf("You chose %s, I chose %s. You lose!", userChoice, botChoice))
		updateHighscore(msg.From.ID, 0, 1, 0)
	default:
		sendStickerAndMessage(msg.Chat.ID, "CAACAgIAAxkBAUnf02YJcCup7gIIO5DMBND1PFZ3seDUAAIbAANd6qsinB_Cwhpp6Uo0BA", fmt.Sprintf("You chose %s, I chose %s. It's a tie!", userChoice, botChoice))
		updateHighscore(msg.From.ID, 0, 0, 1)
	}
}

// Main function to start the bot
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initRedisClient()

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN not found in environment variables")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	startMainLoop(bot)
}

// StartMainLoop starts the main loop to handle updates
func startMainLoop(bot *tgbotapi.BotAPI) {
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
