package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/ports"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/utils"
)

type GreetingsHandler struct {
	svc ports.GreetingsService
}

func (h *GreetingsHandler) MemberRemoveHandlerFunc(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	err := h.svc.RemoveMember(m.User.ID, m.GuildID)
	utils.LogEvent(err)
}

func (h *GreetingsHandler) MemberAddHandlerFunc(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	var err domain.Error

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

	err = h.svc.WelcomeMember(guild.ID, m)
	utils.LogEvent(err)

	err = h.svc.AddMember(member)
	utils.LogEvent(err)

}

func NewGreetingsHandler(svc ports.GreetingsService) *GreetingsHandler {
	return &GreetingsHandler{svc: svc}
}
