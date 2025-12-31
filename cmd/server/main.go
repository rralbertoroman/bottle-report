package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rralbertoroman/bottle-report/internal/app"
	"github.com/rralbertoroman/bottle-report/internal/config"
	"github.com/rralbertoroman/bottle-report/internal/db"
	"github.com/rralbertoroman/bottle-report/internal/handlers"
)

func main() {

	cfg := config.Load()

	database := db.Open(cfg)

	app := &app.App{
		DB: database,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health(app))
	mux.HandleFunc("/messages", handlers.MessagesHandler(app))
	
	server := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("HTTP server listening on %s", cfg.HTTPAddr)
	log.Fatal(server.ListenAndServe())
	// messaging.InitRoutes()
	// fmt.Println("Bottle Report")
	// log.Println("Server started on port 8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}