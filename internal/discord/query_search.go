package discord

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
	"vinted-bidder/internal/search"
)

var lastFetchedItemIDForSearch string

func handleSearchWithQuery(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := trimMentionString(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		return
	}
	fmt.Println(msg)
	var queryParams search.QueryParams
	if strings.Contains(msg, "szukaj") {
		err = json.NewDecoder(strings.NewReader(msg[len("szukaj"):])).Decode(&queryParams)
		if err != nil {
			log.Printf("error decoding query params: %v", err)
		}
		log.Printf("query params: %+v", queryParams)
	}
}

func handleSearch(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := trimMentionString(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.Contains(msg, "szukaj") {
		var searchText string
		if strings.Contains(m.Content, "stone island") {
			searchText = "stone island"
		}

		tool, err := search.New()
		if err != nil {
			fmt.Println(err)
			return
		}

		go func() {
			for {
				response, err := tool.Search(search.QueryParams{
					SearchText: searchText,
				})
				if err != nil {
					fmt.Println(err)
					return
				}

				if len(response.Items) > 0 {
					currentItemID := response.Items[0].Url // Assuming Url is a unique identifier for an item
					if currentItemID != lastFetchedItemIDForSearch {
						lastFetchedItemIDForSearch = currentItemID

						imageUrl := response.Items[0].Photo.Url
						message := response.Items[0].Title + " " + response.Items[0].Price + "PLN " + imageUrl + " " + response.Items[0].SizeTitle + " " + response.Items[0].Url
						_, err := s.ChannelMessageSend(m.ChannelID, message)
						if err != nil {
							fmt.Println(err)
						}
					}
				} else {
					s.ChannelMessageSend(m.ChannelID, "No items found.")
				}
				time.Sleep(5 * time.Second)
			}
		}()

	}
}
