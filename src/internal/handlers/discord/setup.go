package discord

import (
	"fmt"

	"log/slog"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/ports"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/utils"
)

type SetupHandler struct {
	svc ports.SetupService
}

func (h *SetupHandler) SetupHandlerFunc(s *discordgo.Session, m *discordgo.GuildCreate) {
	ok, err := h.svc.IsGuildExists(m.Guild.ID)
	if err.Error != nil {
		slog.Error(err.Error.Error())
	}

	if ok {
		slog.Info(fmt.Sprintf("Guild already exists: %s", m.Guild.ID))
		return
	}

	var settingsID = primitive.NewObjectID().Hex()

	settings := domain.GuildSettings{
		ID: settingsID,
		Welcome: domain.GuildSettingsWelcome{
			Enabled: false,
		},
	}

	guild := domain.Guild{
		ID:         m.Guild.ID,
		Name:       m.Guild.Name,
		OwnerID:    m.Guild.OwnerID,
		SettingsID: settingsID,
	}

	var members []domain.Member

	for _, member := range m.Guild.Members {
		members = append(members, domain.Member{
			ID:            primitive.NewObjectID().Hex(),
			UserID:        member.User.ID,
			GuildID:       member.GuildID,
			Username:      member.User.Username,
			Discriminator: member.User.Discriminator,
			Nickname:      member.Nick,
			AvatarURL:     member.AvatarURL("256"),
			Roles:         member.Roles,
			JoinedAt:      member.JoinedAt,
		})
	}

	var channels []domain.Channel

	for _, channel := range m.Channels {
		channels = append(channels, domain.Channel{
			ID:        primitive.NewObjectID().Hex(),
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
			ID:       primitive.NewObjectID().Hex(),
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

	if err := h.svc.LoadGuild(guild); err.Error != nil {
		slog.Error(err.Error.Error())
	}
	if err := h.svc.LoadSettings(settings); err.Error != nil {
		slog.Error(err.Error.Error())
	}
	if err := h.svc.LoadMembers(members); err.Error != nil {
		slog.Error(err.Error.Error())
	}
	if err := h.svc.LoadChannels(channels); err.Error != nil {
		slog.Error(err.Error.Error())
	}
	if err := h.svc.LoadRoles(roles); err.Error != nil {
		slog.Error(err.Error.Error())
	}

}

func NewSetupHandler(svc ports.SetupService) *SetupHandler {
	return &SetupHandler{svc: svc}
}
