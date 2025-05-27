package db

import (
	"fmt"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/pgtype"
	"os"
)

// get the env variable, but panic if it is an empty string
func envP(s string) string {
	val := os.Getenv(s)
	if val == "" {
		err := fmt.Sprintf("empty string for variable: %s\n", s)
		panic(err)
	}
	return val
}

// get db config variables from env
// panic if any are ""
func connString() string {
	host := envP("DB_HOST")
	port := envP("DB_PORT")
	user := envP("DB_USERNAME")
	password := envP("DB_PASSWORD")
	dbname := envP("DB_NAME")

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Printf("connStr: %s\n", str)
	return str

}

func ErrorCode(err error) (string, bool) {
	pgerr, ok := err.(*pgconn.PgError)
	if !ok {
		return "", false
	}
	return pgerr.Code, true
}
