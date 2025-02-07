package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/domain"
	ports "github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/ports"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SetupHandler struct {
	svc ports.SetupService
}

func (h *SetupHandler) SetupHandlerFunc(s *discordgo.Session, m *discordgo.GuildCreate) {
	guild := domain.Guild{
		ID:      primitive.NewObjectID(),
		GuildID: m.Guild.ID,
		Name:    m.Guild.Name,
		OwnerID: m.Guild.OwnerID,
		Settings: domain.Settings{
			Welcome: domain.WelcomeSettings{
				Enabled: false,
			},
		},
	}

	var members []domain.Member

	for _, member := range m.Guild.Members {
		members = append(members, domain.Member{
			ID:            primitive.NewObjectID(),
			UserID:        member.User.ID,
			GuildID:       member.GuildID,
			Username:      member.User.Username,
			Discriminator: member.User.Discriminator,
			Nickname:      member.Nick,
			Roles:         member.Roles,
			JoinedAt:      member.JoinedAt,
		})
	}

	var channels []domain.Channel

	for _, channel := range m.Channels {
		channels = append(channels, domain.Channel{
			ID:        primitive.NewObjectID(),
			ChannelID: channel.ID,
			GuildID:   channel.GuildID,
			Name:      channel.Name,
			Type:      utils.GetChannelType(channel.Type),
			ParentID:  channel.ParentID,
			Position:  channel.Position,
		})
	}

	var roles []domain.Role

	for _, role := range m.Roles {
		roles = append(roles, domain.Role{
			ID:       primitive.NewObjectID(),
			RoleID:   role.ID,
			Name:     role.Name,
			GuildID:  m.Guild.ID,
			Position: role.Position,
			Color: func() string {
				if role.Color == 0 {
					return ""
				} else {
					return fmt.Sprintf("%x", role.Color)
				}
			}(),
			Managed: role.Managed,
		})
	}

	h.svc.LoadServer(guild)
	h.svc.LoadMembers(members)
	h.svc.LoadChannels(channels)
	h.svc.LoadRoles(roles)

}

func NewSetupHandler(svc ports.SetupService) *SetupHandler {
	return &SetupHandler{svc: svc}
}
