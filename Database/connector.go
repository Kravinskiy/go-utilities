package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"os"
	"github.com/Kravinskiy/go-utilities/Functions"
)

// Backward compatibility: providing connection params as command line arguments
func GetConnection(username *string, password *string, dbHost *string, port *string, dbName *string) (db *gorm.DB, err error) {
	
	if username == nil || password == nil || dbName == nil || dbHost == nil {
		envUsername := os.Getenv("db_user")
		envPassword := os.Getenv("db_password")
		envDbName := os.Getenv("db_name")
		envDbHost := os.Getenv("db_host")
		envDbPort := os.Getenv("db_port")

		username = &envUsername
		password = &envPassword
		dbName = &envDbName
		dbHost = &envDbHost
		port = &envDbPort
	}

	dbURI := functions.GetMysqlDsn(*username, *password, *dbHost, *port, *dbName)
	return gorm.Open(mysql.Open(dbURI), &gorm.Config{})
}