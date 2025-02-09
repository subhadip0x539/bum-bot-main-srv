package utils

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func GetChannelType(channelType discordgo.ChannelType) string {

	switch channelType {
	case discordgo.ChannelTypeGuildText:
		return "text"
	case discordgo.ChannelTypeGuildVoice:
		return "voice"
	case discordgo.ChannelTypeGuildCategory:
		return "category"
	case discordgo.ChannelTypeGuildNews:
		return "news"
	case discordgo.ChannelTypeGuildStageVoice:
		return "stage"
	default:
		return "unknown"
	}
}

func ParseTemplate(template string, values map[string]string) string {
	for key, value := range values {
		placeholder := "{" + key + "}"
		template = strings.ReplaceAll(template, placeholder, value)
	}

	return template
}
