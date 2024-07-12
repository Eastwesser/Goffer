package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

type Highscore struct {
	UserID int64
	Wins   int
	Losses int
	Draws  int
}

func InitRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	pong, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Printf("Redis connected successfully: %s", pong)
}

func UpdateHighscore(userID int64, wins, losses, draws int) error {
	ctx := context.Background()
	key := fmt.Sprintf("highscores:%d", userID)
	err := Client.HSet(ctx, key, map[string]interface{}{
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

func GetHighscore(userID int64) (Highscore, error) {
	ctx := context.Background()
	key := fmt.Sprintf("highscores:%d", userID)
	data, err := Client.HGetAll(ctx, key).Result()
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
