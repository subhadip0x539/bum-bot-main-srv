package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WelcomeSettingsMessage struct {
	Type    string      `bson:"type" json:"type"`
	Content interface{} `bson:"content" json:"content"`
}

type WelcomeSettings struct {
	Enabled   bool                    `bson:"enabled" json:"enabled"`
	ChannelID string                  `bson:"channel_id,omitempty" json:"channel_id"`
	Message   *WelcomeSettingsMessage `bson:"message,omitempty" json:"message"`
}

type Settings struct {
	Welcome WelcomeSettings `bson:"welcome" json:"welcome"`
}

type Guild struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	GuildID  string             `bson:"guild_id" json:"guild_id"`
	Name     string             `bson:"name" json:"name"`
	OwnerID  string             `bson:"owner_id" json:"owner_id"`
	Settings Settings           `bson:"settings" json:"settings"`
}

type Member struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	UserID        string             `bson:"user_id" json:"user_id"`
	GuildID       string             `bson:"guild_id" json:"guild_id"`
	Username      string             `bson:"username" json:"username"`
	Discriminator string             `bson:"discriminator" json:"discriminator"`
	Nickname      string             `bson:"nickname" json:"nickname"`
	Roles         []string           `bson:"roles" json:"roles"`
	JoinedAt      time.Time          `bson:"joined_at" json:"joined_at"`
}

type Channel struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	ChannelID string             `bson:"channel_id" json:"channel_id"`
	GuildID   string             `bson:"guild_id" json:"guild_id"`
	Name      string             `bson:"name" json:"name"`
	Type      string             `bson:"type" json:"type"`
	ParentID  string             `bson:"parent_id" json:"parent_id"`
	Position  int                `bson:"position" json:"position"`
}

type Role struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	RoleID   string             `bson:"role_id" json:"role_id"`
	Name     string             `bson:"name" json:"name"`
	GuildID  string             `bson:"guild_id" json:"guild_id"`
	Position int                `bson:"position" json:"position"`
	Color    string             `bson:"color,omitempty" json:"color"`
	Managed  bool               `bson:"managed" json:"managed"`
}
