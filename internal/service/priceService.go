// Package service provides a set of functions, which include business-logic in it
package service

import (
	"math/rand"

	"github.com/eugenshima/fake-price-provider/internal/model"
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
	PriceStreaming(price []*model.Share)
}

// GeneratePrice function is an infinite loop, which call PriceStreaming function
func (priceService *PriceService) GeneratePrice(constPrices []*model.Share) {
	for {
		share := constPrices
		priceService.rps.PriceStreaming(share)
		for i := 0; i < len(constPrices); i++ {
			price, err := DecimalCalculations(share[i].SharePrice, GenerateRandomFloat())
			if err != nil {
				logrus.Errorf("Generator: %v", err)
			}

			share[i].SharePrice = price.InexactFloat64()
		}
	}
}

// DecimalCalculations returns new price decimal
func DecimalCalculations(price float64, delta float64) (decimal.Decimal, error) {
	deltaPrice := decimal.NewFromFloatWithExponent(delta, -2)
	decPrice := decimal.NewFromFloat(price)

	if rand.Intn(2) == 1 || deltaPrice.GreaterThanOrEqual(decPrice) {
		decPrice = decPrice.Add(deltaPrice)
	} else {
		decPrice = decPrice.Sub(deltaPrice)
	}
	return decPrice, nil
}

// GenerateRandomFloat generates a random float between 0 and 1
func GenerateRandomFloat() float64 {
	randomFloat := rand.Float64() + 1
	return randomFloat
}
