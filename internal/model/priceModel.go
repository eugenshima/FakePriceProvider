// Package model of our entity
package model

// Share struct represents a model for shares
type Share struct {
	SharePrice string `json:"price"`
	ShareName  string `json:"share_name"`
}
