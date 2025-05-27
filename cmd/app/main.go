package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/jmarren/go-web/internal/server"
	"github.com/joho/godotenv"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	// handle flags
	var environment string
	flag.StringVar(&environment, "env", "dev", "application environment")
	flag.Parse()
	os.Setenv("env", environment)

	// load environment variables
	if environment == "dev" {
		godotenv.Load(".local.env")
	} else {
		godotenv.Load(".prod.env")
	}

	defer cancel()
	return server.New(ctx)
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
