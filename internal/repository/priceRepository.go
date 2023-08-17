package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/eugenshima/FakePriceProvider/internal/model"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type PriceRepository struct {
	redisClient *redis.Client
}

func NewPriceRepository(redisClient *redis.Client) *PriceRepository {
	return &PriceRepository{redisClient: redisClient}
}

func (repo *PriceRepository) PriceStreaming(price *model.Share) {
	identificator := 0
	id := strconv.FormatInt(time.Now().Unix(), 10)
	payload := map[string]interface{}{
		"timestamp": id,
		"price":     decimal.Decimal.String(price.SharePrice),
		"name":      price.ShareName,
	}

	identificator++
	id = id + "-" + strconv.Itoa(identificator)

	err := repo.redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "testStream",
		MaxLen: 0,
		ID:     id,
		Values: payload,
	}).Err()
	if err != nil {
		fmt.Println("Error adding message to Redis Stream:", err)
	}
	const TTL = 2
	time.Sleep(TTL * time.Second)

}
