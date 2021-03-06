package models;

import (
	"gorm.io/gorm"
)

type TickerSymbol struct {
	gorm.Model
	Currency string `json:"currency"`
	Description string `json:"description"`
	Symbol string `json:"symbol"`
	SecurityType string `json:"securityType"`
	StockExchange string `json:"stockExchange"`
	Active bool `json:"active"`
}