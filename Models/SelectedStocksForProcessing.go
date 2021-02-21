package models;

import (
	"gorm.io/gorm"
)

type SelectedStocksForProcessing struct {
	gorm.Model
	Id int
	Ticker_symbol_id int
}