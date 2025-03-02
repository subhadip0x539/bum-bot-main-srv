package handlers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bwmarrin/discordgo"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/ports"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/pkg/logger"
)

type GreetingsHandler struct {
	svc ports.GreetingsService
}

func (h *GreetingsHandler) MemberRemoveHandlerFunc(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	err := h.svc.RemoveMember(m.User.ID, m.GuildID)
	if err != nil {
		logger.Error(err.Message)
		return
	}
}

func (h *GreetingsHandler) MemberAddHandlerFunc(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	guild, _ := s.Guild(m.GuildID)

	member := domain.Member{
		ID:            primitive.NewObjectID().Hex(),
		UserID:        m.User.ID,
		GuildID:       m.GuildID,
		Username:      m.User.Username,
		Discriminator: m.User.Discriminator,
		Nickname:      m.Nick,
		AvatarURL:     m.AvatarURL("256"),
		Roles:         m.Roles,
		Bot:           m.User.Bot,
		JoinedAt:      m.JoinedAt,
	}

	if err := h.svc.WelcomeMember(guild.ID, m); err != nil {
		logger.Error(err.Message)
		return
	}

	if err := h.svc.AddMember(member); err != nil {
		logger.Error(err.Message)
		return
	}

}

func NewGreetingsHandler(svc ports.GreetingsService) *GreetingsHandler {
	return &GreetingsHandler{svc: svc}
}
