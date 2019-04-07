package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
)

// GetMySQLDB connection
func GetMySQLDB() *sql.DB {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_SCHEMA")

	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db := createConnectionMySQL(desc)

	return db
}

func createConnectionMySQL(desc string) *sql.DB {
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
