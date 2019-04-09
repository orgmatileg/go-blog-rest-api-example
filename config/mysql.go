package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
)

var db *sql.DB

// InitConnectionDB connection
func InitConnectionDB() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_SCHEMA")

	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db = createConnectionMySQL(desc)
}

func GetMySQLDB() *sql.DB {
	return db
}

func createConnectionMySQL(desc string) *sql.DB {

	fmt.Println(desc)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", desc, val.Encode())
	db, err := sql.Open(`mysql`, dsn)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}
