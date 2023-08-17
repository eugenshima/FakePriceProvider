// Package service provides a set of functions, which include business-logic in it
package service

import (
	"fmt"
	"math/rand"

	"github.com/eugenshima/FakePriceProvider/internal/model"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

// PriceService represents a PriceProvider
type PriceService struct {
	rps PriceRepository
}

// NewPriceService creates a new PriceService
func NewPriceService(rps PriceRepository) *PriceService {
	return &PriceService{rps: rps}
}

// PriceRepository interface represents a repository methods
type PriceRepository interface {
	PriceStreaming(price []model.Share)
}

// GeneratePrice function is an infinite loop, which call PriceStreaming function
func (priceService *PriceService) GeneratePrice(constPrices []model.Share) {
	for {
		share := constPrices
		priceService.rps.PriceStreaming(share)
		for i := 0; i < 10; i++ {
			price, err := DecimalCalculations(share[i].SharePrice, GenerateRandomFloat())
			if err != nil {
				logrus.Errorf("Generator: %v", err)
			}

			share[i].SharePrice = price.String()
		}
	}
}

// DecimalCalculations returns new price decimal
func DecimalCalculations(price string, delta float64) (decimal.Decimal, error) {
	deltaPrice := decimal.NewFromFloatWithExponent(delta, -2)
	decPrice, err := decimal.NewFromString(price)
	if err != nil {
		return decimal.Zero, fmt.Errorf("NewFromString: %v", err)
	}

	if rand.Intn(2) == 1 || deltaPrice.GreaterThanOrEqual(decPrice) {
		decPrice = decPrice.Add(deltaPrice)
	} else {
		decPrice = decPrice.Sub(deltaPrice)
	}
	return decPrice, nil
}

// GenerateRandomFloat generates a random float between 0 and 1
func GenerateRandomFloat() float64 {
	randomFloat := rand.Float64()
	return randomFloat
}
