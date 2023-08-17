package service

import (
	"fmt"
	"math/rand"

	"github.com/eugenshima/FakePriceProvider/internal/model"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type PriceService struct {
	rps PriceRepository
}

func NewPriceService(rps PriceRepository) *PriceService {
	return &PriceService{rps: rps}
}

type PriceRepository interface {
	PriceStreaming(price *model.Share)
}

func (priceService *PriceService) GeneratePrice() {
	price, err := Generator(GenerateRandomFloat())
	if err != nil {
		logrus.Errorf("Generator: %v", err)
	}
	share := &model.Share{
		SharePrice: price,
		ShareName:  "Apple",
	}
	priceService.rps.PriceStreaming(share)
}

func Generator(sharePrice float64) (decimal.Decimal, error) {
	price := decimal.NewFromFloatWithExponent(sharePrice, -2)
	return price, nil
}

func GenerateRandomFloat() float64 {
	randomFloat := rand.Float64()
	fmt.Println(randomFloat)
	return randomFloat
}
