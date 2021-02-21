package types

type Quote struct {
	C float32 // Current price
	H float32 // High price of the day
	L float32 // Low price of the day
	O float32 // Open price of the day
	Pc float32 // Previous close price
}
