package handlers

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bwmarrin/discordgo"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/ports"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/utils"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/pkg/logger"
)

type SetupHandler struct {
	svc ports.SetupService
}

func (h *SetupHandler) SetupHandlerFunc(s *discordgo.Session, m *discordgo.GuildCreate) {
	exists, err := h.svc.IsGuildExists(m.Guild.ID)
	if err != nil {
		logger.Error(err.Message)
		return
	}
	if exists {
		logger.Info(fmt.Sprintf("Guild already exists with id {%s}", m.Guild.ID))
		return
	}

	plugins, err := h.svc.GetPlugins()
	if err != nil {
		logger.Error(err.Message)
		return
	}

	var settingsID = primitive.NewObjectID().Hex()

	settings := domain.GuildSettings{
		ID: settingsID,
		Plugins: func() (g []domain.GuildSettingsPlugin) {
			for _, v := range plugins {
				g = append(g, domain.GuildSettingsPlugin{ID: v.ID, Name: v.Name, Enabled: false})
			}
			return g
		}(),
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
			Bot:           member.User.Bot,
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
			Color:    role.Color,
			Managed:  role.Managed,
		})
	}

	err = h.svc.LoadGuild(guild)
	if err != nil {
		logger.Error(err.Message)
		return
	}

	err = h.svc.LoadSettings(settings)
	if err != nil {
		logger.Error(err.Message)
		return
	}

	err = h.svc.LoadMembers(members)
	if err != nil {
		logger.Error(err.Message)
		return
	}

	err = h.svc.LoadChannels(channels)
	if err != nil {
		logger.Error(err.Message)
		return
	}

	err = h.svc.LoadRoles(roles)
	if err != nil {
		logger.Error(err.Message)
		return
	}
}

func NewSetupHandler(svc ports.SetupService) *SetupHandler {
	return &SetupHandler{svc: svc}
}
