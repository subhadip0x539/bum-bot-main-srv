package services

import (
	"github.com/bwmarrin/discordgo"
	ports "github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/ports"
)

type WelcomeServiceImpl struct {
	repo ports.DiscordRepo
}

func (s *WelcomeServiceImpl) GreetUser(guildID string, embed *discordgo.MessageEmbed) error {
	channel := s.repo.FindChannel(guildID, "general", discordgo.ChannelTypeGuildText)
	s.repo.SendEmbed(channel.ID, embed)
	return nil
}

func NewWelcomeService(repo ports.DiscordRepo) *WelcomeServiceImpl {
	return &WelcomeServiceImpl{repo: repo}
}
