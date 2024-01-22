package discord

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"regexp"
)

// trimMentionString removes the mention string from the message content leaving only the message content
// e.g <@123456789> hello -> hello
func trimMentionString(m *discordgo.MessageCreate) (string, error) {
	re, err := regexp.Compile("<@\\d+>\\s*")
	if err != nil {
		return "", errors.New("error compiling regex")
	}
	result := re.ReplaceAllString(m.Content, "")
	return result, nil
}
