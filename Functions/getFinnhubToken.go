package functions

import (
	"os"
)

func GetFinnhubToken() (string, error) {
	return os.Getenv("finnhub_token"), nil
}