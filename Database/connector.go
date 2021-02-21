package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"github.com/Kravinskiy/go-utilities/Functions"
)

// Backward compatibility: providing connection params as command line arguments
func GetConnection(username *string, password *string, dbName *string, dbHost *string, port *string) (db *gorm.DB, err error) {
	
	if username == nil || password == nil || dbName == nil || dbHost == nil {
		e := godotenv.Load()
		if e != nil {
			fmt.Println(e)
		}

		envUsername := os.Getenv("db_user")
		envPassword := os.Getenv("db_password")
		envDbName := os.Getenv("db_name")
		envDbHost := os.Getenv("db_host")

		username = &envUsername
		password = &envPassword
		dbName = &envDbName
		dbHost = &envDbHost
	}

	dbURI := functions.GetMysqlDsn(*username, *password, *dbHost, *port, *dbName)
	return gorm.Open(mysql.Open(dbURI), &gorm.Config{})
}