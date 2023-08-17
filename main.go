package main

import (
	"fmt"

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

	fmt.Println("Connected to redis!")
	rdb := redis.NewClient(opt)
	return rdb, nil
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Errorf("Error extracting env variables: %v", err)
		return
	}
	client, err := NewRedis(cfg.RedisConnectionString) //TODO: create redis stream
	if err != nil {
		logrus.WithFields(logrus.Fields{"str": cfg.RedisConnectionString}).Errorf("NewRedis: %v", err)
	}
	rps := repository.NewPriceRepository(client)
	ps := service.NewPriceService(rps)

	ps.GeneratePrice([]string{"2", "5", "10", "20", "60", "100", "120", "150", "200", "1000"})

}
