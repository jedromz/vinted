package discord

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"vinted-bidder/internal/client"
	"vinted-bidder/internal/search"
)

const (
	startMessagePrefix = "<@1198201183504957481> start "
	pollInterval       = 5 * time.Second // Interval for polling items
)

func onBotStart(s *discordgo.Session, m *discordgo.MessageCreate) {
	params, err := parseSearchStartParams(m.Content)
	if err != nil {
		log.Printf("Failed to parse start params: %v\n", err)
		return
	}

	tool, err := search.New()
	if err != nil {
		log.Printf("Failed to create search tool: %v\n", err)
		return
	}

	go pollItems(s, m.ChannelID, tool, params)
}

func parseSearchStartParams(message string) (client.IQueryParams, error) {
	var params search.QueryParams
	log.Println(message[len(startMessagePrefix):])
	err := json.NewDecoder(strings.NewReader(message[len(startMessagePrefix):])).Decode(&params)
	return &params, err
}

// pollItems continuously polls for new items and sends messages to a Discord channel.
func pollItems(s *discordgo.Session, channelID string, tool *search.Tool, params client.IQueryParams) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	lastFetchedItemID := ""

	for range ticker.C {
		checkAndSendMessage(s, channelID, tool, params, &lastFetchedItemID)
	}
}

// checkAndSendMessage fetches items and sends a message if a new item is found.
func checkAndSendMessage(s *discordgo.Session, channelID string, tool *search.Tool, params client.IQueryParams, lastFetchedItemID *string) {
	itemsResponse, err := tool.Search(params)
	if err != nil {
		log.Printf("Error searching items: %v\n", err)
		return
	}

	if len(itemsResponse.Items) > 0 {
		currentItemID := itemsResponse.Items[0].Url
		if currentItemID != *lastFetchedItemID {
			*lastFetchedItemID = currentItemID
			sendNewItemMessage(s, channelID, itemsResponse.Items[0])
		} else {
			log.Println("No new items found.")
		}
	} else {
		log.Println("No items found.")
	}
}

// sendNewItemMessage formats and sends a message for a new item.
func sendNewItemMessage(s *discordgo.Session, channelID string, item client.Item) {
	message := fmt.Sprintf(
		"**Title:** %s\n**Price:** %s PLN\n**Photo:** %s\n**Size:** %s\n**URL:** %s",
		item.Title,
		item.Price,
		item.Photo.Url,
		item.SizeTitle,
		item.Url,
	)
	s.ChannelMessageSend(channelID, message)
}
