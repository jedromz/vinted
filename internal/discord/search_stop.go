package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (b *Bot) onBotStop(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Assuming you use the channelID as the process name/key
	err := b.Manager.StopProcess(context.Background(), m.ChannelID)
	if err != nil {
		log.Printf("Failed to stop process: %v\n", err)
	}
}
