package main

import (
	"fmt"
	"time"

	"github.com/eugenshima/FakePriceProvider/internal/config"
	"github.com/eugenshima/FakePriceProvider/internal/repository"
	"github.com/eugenshima/FakePriceProvider/internal/service"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

//NewRedis function provides Connection with Redis database
func NewRedis(env string) (*redis.Client, error) {
	opt, err := redis.ParseURL(env)
	if err != nil {
		return nil, fmt.Errorf("error parsing redis: %v", err)
	}

	logrus.Println("Connected to redis!")
	rdb := redis.NewClient(opt)
	return rdb, nil
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error extracting env variables: %v", err)
		return
	}
	client, err := NewRedis(cfg.RedisConnectionString) //TODO: create redis stream
	if err != nil {
		logrus.WithFields(logrus.Fields{"str": cfg.RedisConnectionString}).Errorf("NewRedis: %v", err)
	}
	rps := repository.NewPriceRepository(client)
	ps := service.NewPriceService(rps)
	for {
		ps.GeneratePrice()
		fmt.Println("price generated, waiting...")
		time.Sleep(3 * time.Second)
	}

}
