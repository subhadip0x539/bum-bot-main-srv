package ports

import (
	"github.com/bwmarrin/discordgo"

	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/domain"
)

type DiscordRepo interface {
	FindChannel(guildID, name string, channelType discordgo.ChannelType) *discordgo.Channel
	SendMessage(channelID, message string) error
	SendEmbed(channelID string, embed *discordgo.MessageEmbed) error
}

type WelcomeService interface {
	GreetUser(guild string, event *discordgo.GuildMemberAdd) domain.Error
}

type SetupService interface {
	IsGuildExists(ID string) (bool, domain.Error)
	LoadGuild(guild domain.Guild) domain.Error
	LoadSettings(settings domain.GuildSettings) domain.Error
	LoadMembers(members []domain.Member) domain.Error
	LoadChannels(channels []domain.Channel) domain.Error
	LoadRoles(roles []domain.Role) domain.Error
}
