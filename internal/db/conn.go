package db

import (
	"context"
	"fmt"

	pgconn "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmarren/go-web/internal/db/query"
)

var (
	Dbtx  *pgxpool.Pool
	Query *query.Queries
)

func Init(ctx context.Context) error {
	connStr := connString()

	DbPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Println("error creating pool")
		return err
	}

	Query = query.New(DbPool)

	if Query == nil {
		return fmt.Errorf("error: nil db.Query !")
	}

	fmt.Printf("database initialized successfully\n")
	return nil
}

func ErrorCode(err error) (string, bool) {
	pgerr, ok := err.(*pgconn.PgError)
	if !ok {
		return "", false
	}
	return pgerr.Code, true
}
