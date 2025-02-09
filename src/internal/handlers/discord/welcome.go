package discord

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/ports"
)

type WelcomeHandler struct {
	svc ports.WelcomeService
}

func (h *WelcomeHandler) WelcomeHandlerFunc(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if err := h.svc.GreetUser(guild.ID, m); err.Error != nil {
		slog.Error(err.Error.Error())
		return
	}
}

func NewWelcomeHandler(svc ports.WelcomeService) *WelcomeHandler {
	return &WelcomeHandler{svc: svc}
}
