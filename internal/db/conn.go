package db

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	// "os"

	pgconn "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmarren/go-web/internal/db/query"
	// "github.com/joho/godotenv"
	// _ "github.com/lib/pq"
)

/*
// go:embed schema.sql
var initQueryProd string

// /go:embed sql/restart_db.sql
var restartDBQuery string
*/

var (
	Dbtx  *pgxpool.Pool
	Query *query.Queries
)

func Init(ctx context.Context, environment string) error {
	fmt.Printf("initializing %s environment\n", environment)

	if environment == "dev" {
		return initDev(ctx)
	}
	// if environment == "prod" {
	// 	return initProd(ctx)
	// }
	return fmt.Errorf("environment not specified")
}

func initDev(ctx context.Context) error {
	connStr := connString()

	Dbtx, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return err
	}

	initQuery, err := getInitQuery()
	if err != nil {
		return fmt.Errorf("error reading schema.sql file: %s", err)
	}

	status, err := Dbtx.Exec(ctx, initQuery)
	if err != nil {
		log.Fatalf("failed to execute init script: %v", err)
	}
	fmt.Printf("status of init exec: %s\n", status)

	Query = query.New(Dbtx)

	if Query == nil {
		return fmt.Errorf("error: nil db.Query !")
	}
	fmt.Printf("db.Query: %v\n", *Query)

	// fmt.Printf("status of testDataQuery exec: %s\n", status)
	fmt.Printf("database initialized successfully\n")
	return nil
}

// func initProd(ctx context.Context) error {
// 	fmt.Println("running production init (initProd)")
// 	// TESTING *****
// 	// initQuery, err := os.ReadFile("/Users/johnmarren//db/sql/schema.sql")
// 	// *********
// 	// initQuery, err := os.ReadFile("~/app/sql/schema.sql")
// 	// if err != nil {
// 	// 	return fmt.Errorf("error reading schema.sql file: %s", err)
// 	// }
//
// 	var err error
// 	Dbtx, err = pgxpool.New(ctx, awssdk.DbDsn)
// 	if err != nil {
// 		return fmt.Errorf("error connection to prod db: %s", err)
// 	}
// 	fmt.Printf("Dbtx: %v\n", Dbtx)
// 	status, err := Dbtx.Exec(ctx, initQueryProd)
// 	if err != nil {
// 		return fmt.Errorf("failed to execute init script: %v", err)
// 	}
// 	fmt.Printf("status of init exec: %s\n", status)
//
// 	Query = New(Dbtx)
// 	if Query == nil {
// 		return fmt.Errorf("error: nil db.Query !")
// 	}
// 	fmt.Printf("db.Query: %v\n", *Query)
// 	if Query == nil {
// 		return fmt.Errorf("error: nil db.Query !")
// 	}
// 	fmt.Printf("db.Query: %v\n", *Query)
//
// 	return nil
// }

func ErrorCode(err error) (string, bool) {
	pgerr, ok := err.(*pgconn.PgError)
	if !ok {
		return "", false
	}
	return pgerr.Code, true
}

// if os.Getenv("resetdb") == "true" {
// 	resetDB, err := os.ReadFile("/home/john-marren/Projects/deepfried/model/sql/restart.sql")
// 	if err != nil {
// 		return fmt.Errorf("error reading restart.sql file: %s", err)
// 	}
// 	_, err = Dbtx.Exec(ctx, string(resetDB))
// 	if err != nil {
// 		log.Fatalf("failed to execute reset script: %v", err)
// 	}
// }
