package ports

import (
	"github.com/bwmarrin/discordgo"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/domain"
)

type DiscordRepo interface {
	FindChannel(guildID, name string, channelType discordgo.ChannelType) *discordgo.Channel
	SendMessage(channelID, message string) error
	SendEmbed(channelID string, embed *discordgo.MessageEmbed) error
}

type GreetingsService interface {
	AddMember(member domain.Member) domain.Error
	RemoveMember(memberID string, guildID string) domain.Error
	WelcomeMember(guildID string, event *discordgo.GuildMemberAdd) domain.Error
	GoodbyeMember(guildID string, event *discordgo.GuildMemberRemove) domain.Error
}

type SetupService interface {
	IsGuildExists(ID string) (bool, domain.Error)
	LoadGuild(guild domain.Guild) domain.Error
	LoadSettings(settings domain.GuildSettings) domain.Error
	LoadMembers(members []domain.Member) domain.Error
	LoadChannels(channels []domain.Channel) domain.Error
	LoadRoles(roles []domain.Role) domain.Error
}
