package functions

import (
	"github.com/joho/godotenv"
	"os"
)

func GetFinnhubToken() (string, error) {
	e := godotenv.Load()
	if e != nil {
		return "", e
	}

	return os.Getenv("finnhub_token"), nil
}