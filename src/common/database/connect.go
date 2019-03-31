package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func Connect() *sql.DB {
	host := os.Getenv("RDS_HOSTNAME")
	user := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	port := os.Getenv("RDS_PORT")
	database := os.Getenv("RDS_DATABASE")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, database))

	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
		return nil
	}

	return db
}
