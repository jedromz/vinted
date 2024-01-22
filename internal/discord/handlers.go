package discord

import "github.com/bwmarrin/discordgo"

type Handler interface {
	Handle(s *discordgo.Session, m interface{})
}
type messageHandler struct {
	handler func(s *discordgo.Session, m *discordgo.MessageCreate)
}

func (h *messageHandler) Handle(s *discordgo.Session, m interface{}) {
	h.handler(s, m.(*discordgo.MessageCreate))
}

type interactionHandler struct {
	handler func(s *discordgo.Session, m *discordgo.InteractionCreate)
}

func (h *interactionHandler) Handle(s *discordgo.Session, m interface{}) {
	h.handler(s, m.(*discordgo.InteractionCreate))
}

func handlerForInterface(handler interface{}) Handler {
	switch handler.(type) {
	case func(s *discordgo.Session, m *discordgo.MessageCreate):
		return &messageHandler{handler.(func(s *discordgo.Session, m *discordgo.MessageCreate))}
	case func(s *discordgo.Session, m *discordgo.InteractionCreate):
		return &interactionHandler{handler.(func(s *discordgo.Session, m *discordgo.InteractionCreate))}
	default:
		return nil
	}
}
