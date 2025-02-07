package ports

import "github.com/bwmarrin/discordgo"

type DiscordRepo interface {
	FindChannel(guildID, name string, channelType discordgo.ChannelType) *discordgo.Channel
	SendMessage(channelID, message string) error
	SendEmbed(channelID string, embed *discordgo.MessageEmbed) error
}

type WelcomeService interface {
	GreetUser(GuildID string, message *discordgo.MessageEmbed) error
}
