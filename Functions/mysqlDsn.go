package functions

import (
	"fmt"
)

func GetMysqlDsn(username string, password string, hostname string, port string, dbname string) string {  
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, dbname)
}