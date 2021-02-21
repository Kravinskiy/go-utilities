package models;

import (
	"gorm.io/gorm"
)

type TickerSymbol struct {
	gorm.Model
	Id int
	Currency string
	Description string
	Symbol string
	Security_type string
	Stock_exchange string
	Active bool
	Created_at string 
}