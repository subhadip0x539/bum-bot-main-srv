package discord

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordClient struct {
	Session *discordgo.Session
}

func (d *DiscordClient) RegisterHandler(handler interface{}) {
	d.Session.AddHandler(handler)
}

func (d *DiscordClient) Start() error {
	err := d.Session.Open()
	if err != nil {
		return err
	}

	return nil
}

func (d *DiscordClient) Stop() error {
	if err := d.Session.Close(); err != nil {
		return err
	}

	return nil
}

func NewDiscordClient(token string) (*DiscordClient, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsAll

	return &DiscordClient{
		Session: session,
	}, nil
}
