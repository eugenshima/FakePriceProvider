package model

import "github.com/shopspring/decimal"

type Share struct {
	SharePrice decimal.Decimal `json:"price"`
	ShareName  string          `json:"share_name"`
}
