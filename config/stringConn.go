package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetStringConn(app string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER_ID")
	pass := os.Getenv("PASS")
	database := os.Getenv("DATABASE")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;app name=%s",
		host, user, pass, database, app)

	return connString
}
