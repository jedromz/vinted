package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type SearchStartParams struct {
	SizeIDs     []int
	Catalog     []int
	MaterialIDs []int
	ColorIDs    []int
	BrandIDs    []int
	PriceFrom   int
	PriceTo     int
	StatusIDs   []int
	Order       string
	Currency    string
	SearchText  string
	Page        int
	PerPage     int
}

func onBotStart(s *discordgo.Session, m *discordgo.MessageCreate) {

	minValues := 1
	maxValues := 2
	brandSelectMenu := discordgo.SelectMenu{
		CustomID:    "brand_select",
		Placeholder: "Select Brands",
		Options: []discordgo.SelectMenuOption{
			{Label: "Adidas", Value: "adidas", Emoji: discordgo.ComponentEmoji{Name: "👟"}},
			{Label: "Prada", Value: "Prada", Emoji: discordgo.ComponentEmoji{Name: "👟"}},
		},
		MinValues: &minValues,
		MaxValues: maxValues,
	}

	// Create action rows to hold the select menus
	actionRow1 := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{&brandSelectMenu},
	}

	// Send a message with the select menus
	_, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content:    "Bot is now running. Choose a configuration option:",
		Components: []discordgo.MessageComponent{&actionRow1},
	})
	if err != nil {
		log.Printf("error sending message: %v", err)
	}
}
func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Check if the interaction is a select menu interaction

	if i.Type == discordgo.InteractionMessageComponent {
		// Check if the custom ID matches
		if i.MessageComponentData().CustomID == "brand_select" {
			selectedBrands := i.MessageComponentData().Values
			response := fmt.Sprintf("You have selected the following brands: %v", selectedBrands)

			// Respond to the interaction
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response,
				},
			})
			if err != nil {
				log.Printf("error responding to select menu interaction: %v", err)
			}
		}
	}
}
