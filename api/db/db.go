package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DBUser = "matteo"
const DBPassword = "passwordMatteo"
const DBName = "detor"
const DBHost = "host.docker.internal"
const DBPort = "3306"

func DbConnection() *gorm.DB {

	var err error

	dns := DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("ERRORE durante la connessione al db: ", err)
		return nil
	}
	return db
}
