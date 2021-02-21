package db

import (
	"fmt"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"github.com/Kravinskiy/go-utilities/Functions"
)

func GetConnection() (db *gorm.DB, err error) {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbURI := Functions.GetMysqlDsn(username, password, hostname, port, dbname)
	return gorm.Open("mysql", dbURI)
}