package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DBHost = "host.docker.internal"

const ContainerName = "apiDockerDB"

func DbConnection() *gorm.DB {
	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}
	DBUser := os.Getenv("MYSQL_USER")
	DBPassword := os.Getenv("MYSQL_PASSWORD")
	DBName := os.Getenv("MYSQL_DATABASE")

	dns := DBUser + ":" + DBPassword + "@tcp(" + ContainerName + ")/" + DBName + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("ERRORE durante la connessione al db: ", err)
		return nil
	}
	return db
}
