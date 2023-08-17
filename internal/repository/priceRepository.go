// Package repository provides functions for interacting with a redis stream
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/eugenshima/FakePriceProvider/internal/model"
	"github.com/redis/go-redis/v9"
)

// PriceRepository represents a repository structure
type PriceRepository struct {
	redisClient *redis.Client
}

// NewPriceRepository creates a new PriceRepository
func NewPriceRepository(redisClient *redis.Client) *PriceRepository {
	return &PriceRepository{redisClient: redisClient}
}

// PriceStreaming function streams all share prices to redis stream
func (repo *PriceRepository) PriceStreaming(price []model.Share) {
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
