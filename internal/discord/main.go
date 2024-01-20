package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	search2 "vinted-bidder/internal/search"
)

func main() {
	token := "MTE5ODIwMTE4MzUwNDk1NzQ4MQ.GWI-NT.ch1IpGuRC9VBgnqucDSS_b7kf6_NG-MEvSPrR4"
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check the message is not sent by a bot
	if m.Author.Bot {
		return
	}
	// Compile the regular expression
	re, err := regexp.Compile("<@\\d+>\\s*")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Replace all occurrences of the pattern with an empty string
	result := re.ReplaceAllString(m.Content, "")
	if result == "!fetch" {
		// Call your function to fetch items from a webpage
		items := fetchItemsFromWebPage()

		// Send the fetched items back as a message
		s.ChannelMessageSend(m.ChannelID, items)
	}
	fmt.Println(result)
	if result == "szukaj" {
		tool, err := search2.New()
		if err != nil {
			fmt.Println(err)
			return
		}

		response, err := tool.Search(search2.QueryParams{})
		if err != nil {
			fmt.Println(err)
			return
		}

		// Check if there are any items in the response
		if len(response.Items) > 0 {
			// Construct the message with the image URL
			imageUrl := response.Items[0].Photo.Url
			message := response.Items[0].Title + " " + response.Items[0].Price + "PLN " + imageUrl + " " + response.Items[0].SizeTitle + " " + response.Items[0].Url

			// Send the message to the Discord channel
			s.ChannelMessageSend(m.ChannelID, message)
		} else {
			// Send a message if no items are found
			s.ChannelMessageSend(m.ChannelID, "No items found.")
		}
	}

}

func fetchItemsFromWebPage() string {
	// Your existing GoLang program logic to fetch items
	// Return the items as a string
	return "Fetched Items"
}
