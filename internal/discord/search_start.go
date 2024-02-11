package discord

import (
	"context"
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
	pollInterval       = 5 * time.Second
)

func (b *Bot) onBotStart(s *discordgo.Session, m *discordgo.MessageCreate) {
	params, err := parseSearchStartParams(m.Content)
	if err != nil {
		log.Printf("Failed to parse start params: %v", err)
		return
	}

	tool, err := search.New()
	if err != nil {
		log.Printf("Failed to create search tool: %v", err)
		return
	}

	lkp := Lookup{
		s:         s,
		channelID: m.ChannelID,
		tool:      tool,
		params:    params,
		stop:      make(chan struct{}),
	}

	b.Manager.AddProcess(m.ChannelID, &lkp)
	if err := b.Manager.StartProcess(context.Background(), m.ChannelID); err != nil {
		log.Printf("Failed to start process for channel %s: %v", m.ChannelID, err)
	}
}

func parseSearchStartParams(message string) (client.IQueryParams, error) {
	messageContent := strings.TrimPrefix(message, startMessagePrefix)
	var params search.QueryParams
	if err := json.NewDecoder(strings.NewReader(messageContent)).Decode(&params); err != nil {
		return nil, fmt.Errorf("parsing search start params: %w", err)
	}
	return &params, nil
}

type Lookup struct {
	s         *discordgo.Session
	channelID string
	tool      *search.Tool
	params    client.IQueryParams
	stop      chan struct{}
}

func (l *Lookup) Start(ctx context.Context) error {
	go l.pollItems(ctx)
	return nil
}
func (l *Lookup) pollItems(ctx context.Context) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	var lastFetchedItemID string

	for {
		select {
		case <-ctx.Done():
			return
		case <-l.stop:
			return
		case <-ticker.C:
			l.checkAndSendMessage(&lastFetchedItemID)
		}
	}
}
func (l *Lookup) Stop(ctx context.Context) error {
	close(l.stop)
	log.Printf("Lookup stopped for channel: %s", l.channelID)
	return nil
}
func (l *Lookup) checkAndSendMessage(lastFetchedItemID *string) {
	itemsResponse, err := l.tool.Search(l.params)
	if err != nil {
		log.Printf("Error searching items: %v", err)
		return
	}

	if len(itemsResponse.Items) == 0 {
		log.Println("No items found.")
		return
	}

	currentItemID := itemsResponse.Items[0].Url
	if currentItemID == *lastFetchedItemID {
		log.Println("No new items found.")
		return
	}

	*lastFetchedItemID = currentItemID
	sendNewItemMessage(l.s, l.channelID, itemsResponse.Items[0])
}

func sendNewItemMessage(s *discordgo.Session, channelID string, item client.Item) {
	message := fmt.Sprintf("**Title:** %s\n**Price:** %s PLN\n**Photo:** %s\n**Size:** %s\n**URL:** %s",
		item.Title, item.Price, item.Photo.Url, item.SizeTitle, item.Url)
	s.ChannelMessageSend(channelID, message)
}
