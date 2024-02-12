package main

import (
	"log"
	"net/http"
	"os"
	"vinted-bidder/internal/discord"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message indicating everything is okay.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	bot, err := discord.New(os.Getenv("BOT_TOKEN")) // Replace with your actual token
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
	}

	err = bot.Start()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}

	// Set up the HTTP server for health checks.
	http.HandleFunc("/health", healthCheckHandler)
	go func() {
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
			log.Fatalf("Failed to start HTTP server for health checks: %v", err)
		}
	}()

	// Block forever.
	select {}
}
