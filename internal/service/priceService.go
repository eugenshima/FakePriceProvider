package service

import (
	"fmt"
	"math/rand"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	MAX = 1
	MIN = 0
)

type PriceService struct {
	rps PriceRepository
}

func NewPriceService(rps PriceRepository) *PriceService {
	return &PriceService{rps: rps}
}

type PriceRepository interface {
	PriceStreaming(price []string)
}

func (priceService *PriceService) GeneratePrice(constPrices []string) {
	for {
		share := constPrices
		priceService.rps.PriceStreaming(share)
		for i := 0; i < 10; i++ {
			price, err := DecimalCalculations(share[i], GenerateRandomFloat())
			if err != nil {
				logrus.Errorf("Generator: %v", err)
			}

			share[i] = price.String()
		}
	}

}

func DecimalCalculations(price string, delta float64) (decimal.Decimal, error) {
	decPrice, err := decimal.NewFromString(price)
	if err != nil {
		return decimal.Zero, fmt.Errorf("NewFromString: %v", err)
	}
	deltaPrice := decimal.NewFromFloat(delta)
	if rand.Intn(2) == 1 && decPrice.GreaterThan(deltaPrice) {
		decPrice = decPrice.Add(deltaPrice)
	} else {
		decPrice = decPrice.Sub(deltaPrice)
	}
	return decPrice, nil
}

func GenerateRandomFloat() float64 {
	randomFloat := rand.Float64()
	return randomFloat
}
