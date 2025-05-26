package db

import (
	"fmt"
	"os"
)

func connString() string {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
