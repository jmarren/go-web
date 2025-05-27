package db

import (
	"fmt"
	"os"
)

func connString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Printf("connStr: %s\n", str)
	return str

}

func getInitQuery() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	queryLocation := wd + "/internal/db/schema.sql"

	initQuery, err := os.ReadFile(queryLocation)

	return string(initQuery), err
}
