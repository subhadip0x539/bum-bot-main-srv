package repositories

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordRepoImpl struct {
	session *discordgo.Session
}

func (r *DiscordRepoImpl) FindChannel(guildID, name string, channelType discordgo.ChannelType) *discordgo.Channel {
	channels, _ := r.session.GuildChannels(guildID)

	for _, channel := range channels {
		if channel.Name == name && channel.Type == channelType {
			return channel
		}
	}
	return nil
}

func (r *DiscordRepoImpl) SendMessage(channelID, message string) error {
	if _, err := r.session.ChannelMessageSend(channelID, message); err != nil {
		return err
	}

	return nil
}

func (r *DiscordRepoImpl) SendEmbed(channelID string, embed *discordgo.MessageEmbed) error {
	if _, err := r.session.ChannelMessageSendEmbed(channelID, embed); err != nil {
		return err
	}

	return nil
}

func NewDiscordRepo(session *discordgo.Session) *DiscordRepoImpl {
	return &DiscordRepoImpl{session: session}
}
