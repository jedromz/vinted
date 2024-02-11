package discord

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"vinted-bidder/internal/manager"
)

type Bot struct {
	Token    string
	Dg       *discordgo.Session
	Handlers map[string]Handler
	Manager  *manager.Manager
}

type EventHandler interface {
}

func New(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	bot := &Bot{Token: token, Dg: dg, Handlers: make(map[string]Handler)}
	bot.Manager = manager.New()
	bot.AddHandler("start", handlerForInterface(bot.onBotStart))
	bot.AddHandler("stop", handlerForInterface(bot.onBotStop))
	bot.Dg.AddHandler(bot.globalHandler)
	return bot, nil
}

func (b *Bot) globalHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	msg, err := trimMentionString(m)
	if err != nil {
		return
	}
	cmd, err := command(msg)
	if err != nil {
		return
	}
	if handler, ok := b.Handlers[cmd]; ok {
		handler.Handle(s, m)
	} else {
		log.Printf("no handler for command: %v", cmd)
	}
}
func command(message string) (string, error) {
	parts := strings.Fields(message)
	if len(parts) > 0 {
		return parts[0], nil
	}
	return "", errors.New("no command")
}

func (b *Bot) Start() error {
	if err := b.Dg.Open(); err != nil {
		return err
	}
	log.Printf("Bot is now running. Press CTRL+C to exit.")
	return nil
}

func (b *Bot) AddHandler(command string, handler Handler) {
	b.Handlers[command] = handler
}
