package main

import (
	"github.com/yousysadmin/nsq-auth/app"
	"github.com/yousysadmin/nsq-auth/app/handlers"
	"log"
	"net/http"
	"os"
)

var Version = "development"

func init() {
	log.Printf("NSQ-Auth %s is starting...", Version)
	// Load configs file
	_, err := app.AppConfig.Load("nsqauth.yaml")
	if err != nil {
		log.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}

func main() {
	// Authenticate endpoint
	http.HandleFunc("/auth", handlers.AuthHandler)

	// Healthcheck endpoint
	http.HandleFunc("/ping", handlers.PingHandler)

	// Start HTTP server
	log.Printf("INFO: started on address %s, identities list size: %d", app.AppConfig.BindAddr, len(app.AppConfig.Identities))
	if err := http.ListenAndServe(app.AppConfig.BindAddr, nil); err != nil {
		log.Printf("ERROR: %s", err)
	}
}
