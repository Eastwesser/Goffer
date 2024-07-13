package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// Client is the global Redis client
var Client *redis.Client

// Highscore represents the highscore data for a user
type Highscore struct {
	UserID int64
	Wins   int
	Losses int
	Draws  int
}

// InitRedisClient initializes the Redis client and connects to the Redis server
func InitRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	ctx := context.Background()
	pong, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Printf("Redis connected successfully: %s", pong)
}

// UpdateHighscore updates the highscore for a user in Redis
func UpdateHighscore(userID int64, wins, losses, draws int) error {
	ctx := context.Background()
	key := fmt.Sprintf("highscores:%d", userID) // Create a key for the user's highscore
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

// GetHighscore retrieves the highscore for a user from Redis
func GetHighscore(userID int64) (Highscore, error) {
	ctx := context.Background()
	key := fmt.Sprintf("highscores:%d", userID) // Create a key for the user's highscore
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
