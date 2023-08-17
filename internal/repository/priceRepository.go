package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type PriceRepository struct {
	redisClient *redis.Client
}

func NewPriceRepository(redisClient *redis.Client) *PriceRepository {
	return &PriceRepository{redisClient: redisClient}
}

func (repo *PriceRepository) PriceStreaming(price []string) {
	identificator := 0
	id := strconv.FormatInt(time.Now().Unix(), 10)
	recordJSON, _ := json.Marshal(price)
	payload := map[string]interface{}{
		"timestamp":       id,
		"GeneratedPrices": recordJSON,
	}

	id = id + "-" + strconv.Itoa(identificator)

	err := repo.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "PriceStreaming",
		MaxLen: 0,
		ID:     id,
		Values: payload,
	}).Err()
	if err != nil {
		fmt.Println("Error adding message to Redis Stream:", err)
	}
	time.Sleep(1 * time.Second)
}
