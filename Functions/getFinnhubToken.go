package Functions

import (
	"github.com/joho/godotenv"
	"os"
)

func getFinnhubToken() string {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	return os.Getenv("finnhub_token")
}