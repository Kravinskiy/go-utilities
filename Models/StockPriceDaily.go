package models;

import (
	"gorm.io/gorm"
)

const StockPriceDailyType = "StockPriceDaily"

type StockPriceDaily struct {
	gorm.Model
	// Ticker Symbol
	TickerSymbolID int `json:"tickerSymbolId"`
	TickerSymbol TickerSymbol `json:"tickerSymbol"`

	// Prices
	OpenPrice float64 `json:"openPrice"`
	HighPrice float64 `json:"highProce"`
	CurrentPrice float64 `json:"currentPrice"`
	PreviousClosePrice float64 `json:"previousClosePrice"`
	LowPrice float64 `json:"lowPrice"`
	RegularMarketChangePercent float64 `json:"regularMarketChangePercent"`
	RegularMarketChange float64 `json:"regularMarketChange"`

	// Volumes
	RegularMarketVolume int `json:"regularMarketVolume"`
	AverageDailyVolume3Month int `json:"averageDailyVolume3Month"`
  	AverageDailyVolume10Day  int `json:"averageDailyVolume10Day"`

	// Valuation Metrics
	TrailingPe float64 `json:"trailingPe"`
	ForwardPe float64 `json:"forwardPe"`
	BookValue float64 `json:"bookValue"`
	PriceToBook float64 `json:"priceToBook"`
	TrailingAnnualDividendYield float64 `json:"trailingAnnualDividendYield"`
	TrailingAnnualDividendRate float64 `json:"trailingAnnualDividendRate"`
	MarketCap int64 `json:"marketCap"`

	// Shares Outstanding
	SharesOutstanding int `json:"sharesOutstanding"`

	// Lows and Highs
	FiftyTwoWeekLow float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekLowChange float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent float64 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekHigh float64 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekHighChange float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent float64 `json:"fiftyTwoWeekHighChangePercent"`

	// Averages
	FiftyDayAverage float64 `json:"FiftyDayAverage"`
	FiftyDayAverageChange float64 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent float64 `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage  float64 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange float64 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
}
  
func (StockPriceDaily) TableName() string {
	// By default it is `stock_price_dailies` - which is wrong. So override that
	return "stock_price_daily"
}