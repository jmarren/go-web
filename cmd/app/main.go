package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/jmarren/go-web/internal/server"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	defer cancel()
	return server.New()
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
