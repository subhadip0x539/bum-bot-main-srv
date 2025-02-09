package domain

import "time"

type GuildSettingsWelcomeMessageType string

const (
	GUILD_SETTINGS_WELCOME_MESSAGE_TYPE_EMBED GuildSettingsWelcomeMessageType = "EMBED"
	GUILD_SETTINGS_WELCOME_MESSAGE_TYPE_TEXT  GuildSettingsWelcomeMessageType = "TEXT"
)

type GuildSettingsWelcomeMessageContent struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Color       string `bson:"color"`
	Image       string `bson:"image"`
	Avatar      bool   `bson:"avatar"`
}

type GuildSettingsWelcomeMessage struct {
	Type    GuildSettingsWelcomeMessageType    `bson:"type" json:"type"`
	Content GuildSettingsWelcomeMessageContent `bson:"content" json:"content"`
}

type GuildSettingsWelcome struct {
	Enabled   bool                         `bson:"enabled" json:"enabled"`
	ChannelID string                       `bson:"channel_id,omitempty" json:"channel_id"`
	Message   *GuildSettingsWelcomeMessage `bson:"message,omitempty" json:"message"`
}

type GuildSettings struct {
	ID      string               `bson:"_id"`
	Welcome GuildSettingsWelcome `bson:"welcome" json:"welcome"`
}

type Guild struct {
	ID         string `bson:"_id" json:"_id"`
	Name       string `bson:"name" json:"name"`
	OwnerID    string `bson:"owner_id" json:"owner_id"`
	SettingsID string `bson:"settings_id" json:"settings_id"`
}

type Member struct {
	ID            string    `bson:"_id" json:"_id"`
	UserID        string    `bson:"user_id" json:"user_id"`
	GuildID       string    `bson:"guild_id" json:"guild_id"`
	Username      string    `bson:"username" json:"username"`
	Discriminator string    `bson:"discriminator" json:"discriminator"`
	Nickname      string    `bson:"nickname" json:"nickname"`
	AvatarURL     string    `bson:"avatar_url"`
	Roles         []string  `bson:"roles" json:"roles"`
	JoinedAt      time.Time `bson:"joined_at" json:"joined_at"`
}

type Channel struct {
	ID        string `bson:"_id" json:"_id"`
	ChannelID string `bson:"channel_id" json:"channel_id"`
	GuildID   string `bson:"guild_id" json:"guild_id"`
	Name      string `bson:"name" json:"name"`
	Type      string `bson:"type" json:"type"`
	ParentID  string `bson:"parent_id" json:"parent_id"`
	Position  int    `bson:"position" json:"position"`
}

type Role struct {
	ID       string `bson:"_id" json:"_id"`
	RoleID   string `bson:"role_id" json:"role_id"`
	Name     string `bson:"name" json:"name"`
	GuildID  string `bson:"guild_id" json:"guild_id"`
	Position int    `bson:"position" json:"position"`
	Color    string `bson:"color,omitempty" json:"color"`
	Managed  bool   `bson:"managed" json:"managed"`
}
