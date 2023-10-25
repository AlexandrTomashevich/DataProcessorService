package main

import (
	"DataProcessorService/app/database/postgres"
	"DataProcessorService/app/internal/api"
	"DataProcessorService/app/internal/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	cfg := config.InitConfig()

	router := mux.NewRouter()
	server := &http.Server{
		Addr:    cfg.ServerHTTP.Address,
		Handler: router,
	}

	dbConnect, err := postgres.NewConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialized database: %s", err)
	}

	s := api.NewServer(dbConnect)

	router.HandleFunc("/api/event", s.Event).Methods("POST")

	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Error starting the server: %s", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	wg.Add(1)

	go func() {
		defer wg.Done()
		<-signals

		if err := dbConnect.Close(); err != nil {
			log.Printf("Error closing the database connection: %s", err)
		}

		if err := server.Shutdown(nil); err != nil {
			log.Fatalf("Error shutting down the server: %s", err)
		}
	}()

	wg.Wait()
}
