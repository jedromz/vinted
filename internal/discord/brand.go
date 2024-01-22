package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
	search2 "vinted-bidder/internal/search"
)

var lastFetchedItemID string

func handleBrand(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := trimMentionString(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.Contains(msg, "stone island") {
		searchText := msg
		t, err := search2.New()
		if err != nil {
			fmt.Println(err)
			return
		}

		go func() {
			for {
				response, err := t.Search(search2.QueryParams{
					SearchText: searchText,
				})
				if err != nil {
					fmt.Println(err)
					time.Sleep(5 * time.Second)
					continue
				}

				currentItemID := response.Items[0].Url // Assuming Url is a unique identifier for an item
				if currentItemID != lastFetchedItemID {
					lastFetchedItemID = currentItemID
					s.ChannelMessageSend(m.ChannelID, response.Items[0].Title+" "+response.Items[0].Price+"PLN "+response.Items[0].Photo.Url+" "+response.Items[0].SizeTitle+" "+response.Items[0].Url)
				} else {
					fmt.Println("No new items found.")
				}

				time.Sleep(5 * time.Second)
			}
		}()
	}
}
