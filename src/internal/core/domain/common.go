package domain

import "time"

type GuildSettingsPluginWelcomeOptionsMessageType string

const (
	GUILD_SETTINGS_WELCOME_MESSAGE_TYPE_EMBED GuildSettingsPluginWelcomeOptionsMessageType = "EMBED"
	GUILD_SETTINGS_WELCOME_MESSAGE_TYPE_TEXT  GuildSettingsPluginWelcomeOptionsMessageType = "TEXT"
)

type GuildSettingsPluginWelcomeOptionsMessageContent struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Color       int    `bson:"color"`
	Image       string `bson:"image"`
	Avatar      bool   `bson:"avatar"`
}

type GuildSettingsPluginWelcomeOptionsMessage struct {
	Type    GuildSettingsPluginWelcomeOptionsMessageType    `bson:"type"`
	Content GuildSettingsPluginWelcomeOptionsMessageContent `bson:"content"`
}

type GuildSettingsPluginWelcomeOptions struct {
	ChannelID string                                    `bson:"channel_id,omitempty"`
	Message   *GuildSettingsPluginWelcomeOptionsMessage `bson:"message,omitempty"`
}

type GuildSettingsServerStatsFields struct {
	MemberCount bool `bson:"member_count"`
}

const (
	StatTotalMembers = "total_members"
)

type GuildSettingsServerStatsChannel struct {
	Enabled   bool   `bson:"enabled"`
	ChannelID string `bson:"channel_id,omitempty"`
	Template  string `bson:"template"`
}

type GuildSettingsServerStatsChannels struct {
	MemberCount GuildSettingsServerStatsChannel `bson:"member_count"`
}

type GuildSettingsServerStats struct {
	Enabled  bool                              `bson:"enabled"`
	ParentID string                            `bson:"parent_id,omitempty"`
	Channels *GuildSettingsServerStatsChannels `bson:"channels,omitempty"`
}

type GuildSettingsPlugin struct {
	ID      string      `bson:"id"`
	Name    string      `bson:"name"`
	Enabled bool        `bson:"enabled"`
	Options interface{} `bson:"options,omitempty"`
}

type GuildSettings struct {
	ID      string                `bson:"_id"`
	Plugins []GuildSettingsPlugin `bson:"plugins"`
}

type Guild struct {
	ID         string `bson:"_id"`
	Name       string `bson:"name"`
	OwnerID    string `bson:"owner_id"`
	SettingsID string `bson:"settings_id"`
}

type Member struct {
	ID            string    `bson:"_id"`
	UserID        string    `bson:"user_id"`
	GuildID       string    `bson:"guild_id"`
	Username      string    `bson:"username"`
	Discriminator string    `bson:"discriminator"`
	Nickname      string    `bson:"nickname"`
	AvatarURL     string    `bson:"avatar_url"`
	Roles         []string  `bson:"roles"`
	Bot           bool      `bson:"bot"`
	JoinedAt      time.Time `bson:"joined_at"`
}

type Channel struct {
	ID        string `bson:"_id"`
	ChannelID string `bson:"channel_id"`
	GuildID   string `bson:"guild_id"`
	Name      string `bson:"name"`
	Type      string `bson:"type"`
	ParentID  string `bson:"parent_id"`
	Position  int    `bson:"position"`
}

type Role struct {
	ID       string `bson:"_id"`
	RoleID   string `bson:"role_id"`
	Name     string `bson:"name"`
	GuildID  string `bson:"guild_id"`
	Position int    `bson:"position"`
	Color    int    `bson:"color,omitempty"`
	Managed  bool   `bson:"managed"`
}

type Plugin struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Label       string `bson:"label"`
	Category    string `bson:"catregory"`
	Description string `bson:"description"`
}
