package discord

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	ports "github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/ports"
)

type WelcomeHandler struct {
	svc ports.WelcomeService
}

func (h *WelcomeHandler) WelcomeHandlerFunc(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		slog.Error("Error fetching guild: " + err.Error())
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: "Welcome to " + guild.Name + "!",
		Description: fmt.Sprintf(
			"Hey %s! Welcome aboard. Check out the rules, introduce yourself, and let's get started!",
			m.Member.Mention(),
		),

		Color: 0xc356fd,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://media.giphy.com/media/26gsi66dCaiuDn5XG/giphy.gif?cid=ecf05e47vwidck6x3v94n7dci9tv8r5rtzn1g2pwz5ujru8m&ep=v1_gifs_related&rid=giphy.gif&ct=g",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: m.AvatarURL("256"),
		},
	}

	h.svc.GreetUser(m.GuildID, embed)
}

func NewWelcomeHandler(svc ports.WelcomeService) *WelcomeHandler {
	return &WelcomeHandler{svc: svc}
}
