// Package main provides entry point to FakePriceProvider
package main

import (
	"fmt"

	"github.com/eugenshima/fake-price-provider/internal/config"
	"github.com/eugenshima/fake-price-provider/internal/model"
	"github.com/eugenshima/fake-price-provider/internal/repository"
	"github.com/eugenshima/fake-price-provider/internal/service"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// NewRedis function provides Connection with Redis database
func NewRedis(env string) (*redis.Client, error) {
	opt, err := redis.ParseURL(env)
	if err != nil {
		return nil, fmt.Errorf("error parsing redis: %v", err)
	}

	fmt.Println("Connected to redis!")
	rdb := redis.NewClient(opt)
	return rdb, nil
}

// main function to run the application
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Errorf("Error extracting env variables: %v", err)
		return
	}
	client, err := NewRedis(cfg.RedisConnectionString)
	if err != nil {
		logrus.WithFields(logrus.Fields{"str": cfg.RedisConnectionString}).Errorf("NewRedis: %v", err)
	}
	rps := repository.NewPriceRepository(client)
	ps := service.NewPriceService(rps)
	start := []*model.Share{
		{
			SharePrice: 1337,
			ShareName:  "Netflix",
		},
		{
			SharePrice: 120,
			ShareName:  "GameStop",
		},
		{
			SharePrice: 229,
			ShareName:  "Spotify",
		},
		{
			SharePrice: 1400,
			ShareName:  "Microsoft",
		},
		{
			SharePrice: 1000,
			ShareName:  "Intel",
		},
	}
	ps.GeneratePrice(start)
}
