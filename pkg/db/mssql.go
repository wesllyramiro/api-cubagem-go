package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/wesllyramiro/api-cubagem-go/config"
)

func Conn() *sql.DB {
	stringConn := config.GetStringConn("api-cubagem-go")

	conn, err := sql.Open("mssql", stringConn)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")

	return conn
}
