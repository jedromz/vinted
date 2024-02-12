package main

import (
	"log"
	"net/http"
	"vinted-bidder/internal/discord"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message indicating everything is okay.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	bot, err := discord.New("MTE5ODIwMTE4MzUwNDk1NzQ4MQ.GWI-NT.ch1IpGuRC9VBgnqucDSS_b7kf6_NG-MEvSPrR4") // Replace with your actual token
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
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Failed to start HTTP server for health checks: %v", err)
		}
	}()

	// Block forever.
	select {}
}
