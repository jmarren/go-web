package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/jmarren/go-web/config"
	"github.com/jmarren/go-web/internal/db"
)

type Server struct {
	*http.Server
}

func New(ctx context.Context) error {

	// create ServeMux
	mux := http.NewServeMux()

	err := db.Init(ctx)
	if err != nil {
		panic(err)
	}

	// add routes
	addRoutes(mux)

	// create server
	s := &Server{
		&http.Server{
			Addr:    config.Port,
			Handler: mux,
		},
	}

	// listen and serve
	go func() {
		log.Printf("listening on %s\n", s.Addr)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	// create wait group and add 1
	var wg sync.WaitGroup
	wg.Add(1)

	// create context
	// ctx := context.Background()

	// spin off routine to handle shutdown
	go func() {
		// defer done
		defer wg.Done()
		// send done to channel
		<-ctx.Done()
		// create shutdown context
		shutdownCtx := context.Background()
		// add timeout
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		// defer canceling context
		defer cancel()
		// shutdown
		if err := s.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil

}

// func NewServer(
// 	logger *Logger
// 	config *Config
// 	commentStore *commentStore
// 	anotherStore *anotherStore
// ) http.Handler {
// 	mux := http.NewServeMux()
// 	addRoutes(
// 		mux,
// 		Logger,
// 		Config,
// 		commentStore,
// 		anotherStore,
// 	)
// 	var handler http.Handler = mux
// 	handler = someMiddleware(handler)
// 	handler = someMiddleware2(handler)
// 	handler = someMiddleware3(handler)
// 	return handler
// }
//
