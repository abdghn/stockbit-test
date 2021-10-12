/*
 * Created on 11/10/21 13.59
 *
 * Copyright (c) 2021 Abdul Ghani Abbasi
 */

package main

import (
	"context"
	"github.com/abdghn/stockbit-test/microservice/internal/handler"
	"github.com/abdghn/stockbit-test/microservice/internal/resource/db"
	"github.com/abdghn/stockbit-test/microservice/internal/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	persistentDB, err := db.NewPersistent("root:stockbit@tcp(127.0.0.1:3306)/stockbit-test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	uc := usecase.New(persistentDB)

	h := handler.New(uc)

	r := mux.NewRouter()
	r.HandleFunc("/health", h.HandleHealthCheck).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/movie", h.HandleSearchMovies).Methods(http.MethodGet)
	r.HandleFunc("/movie/{id}", h.HandleGetMovie).Methods(http.MethodGet)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-sigChan

		cancel()
	}()

	if err := startServer(ctx, ":3000", r); err != nil {
		log.Fatalf("start server failed: %v", err)
	}
}

func startServer(ctx context.Context, addr string, handler http.Handler) error {
	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("starting server failed: %v", err)
		}
	}()

	log.Printf("server is running at %s", srv.Addr)

	// wait for signal to shut down the server gracefully
	<-ctx.Done() // blocking

	log.Printf("shutting down server...")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("an error occurred while shutting down the server: %v", err)
	}

	log.Printf("gracefull shutdown success")
	return nil
}
