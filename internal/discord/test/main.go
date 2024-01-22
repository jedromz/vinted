package main

import (
	"log"
	"vinted-bidder/internal/discord"
)

func main() {
	bot, err := discord.New("MTE5ODIwMTE4MzUwNDk1NzQ4MQ.GWI-NT.ch1IpGuRC9VBgnqucDSS_b7kf6_NG-MEvSPrR4")
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
	}
	err = bot.Start()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}
	select {}
}

func fetchItemsFromWebPage() string {
	// Your existing GoLang program logic to fetch items
	// Return the items as a string
	return "Fetched Items"
}
