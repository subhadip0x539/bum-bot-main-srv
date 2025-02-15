package services

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/ports"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/utils"
)

type GreetingsServiceImpl struct {
	discordRepo ports.DiscordRepo
	mongoRepo   ports.MongoRepo
}

func (s *GreetingsServiceImpl) AddMember(member domain.Member) domain.Error {
	if err := s.mongoRepo.InsertOne("members", member); err != nil {
		return domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message:  fmt.Sprintf("Inserted member successfully with id {%s}", member.ID),
	}
}

func (s *GreetingsServiceImpl) RemoveMember(memberID string, guildID string) domain.Error {
	if err := s.mongoRepo.DeleteOne("members", bson.M{"user_id": memberID, "guild_id": guildID}); err != nil {
		return domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Error:    err,
			Message:  err.Error(),
		}
	}
	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message:  fmt.Sprintf("Deleted member successfully with id {%s}", memberID),
	}
}

func (s *GreetingsServiceImpl) WelcomeMember(guildID string, event *discordgo.GuildMemberAdd) domain.Error {
	pipeline := []bson.M{
		{
			"$match": bson.M{"_id": guildID},
		},
		{
			"$lookup": bson.M{
				"from":         "settings",
				"localField":   "settings_id",
				"foreignField": "_id",
				"as":           "settings",
			},
		},
		{
			"$unwind": "$settings",
		},
		{
			"$project": bson.M{
				"_id":      1,
				"name":     1,
				"settings": 1,
			},
		},
	}

	type Guild struct {
		GuildID   string               `bson:"_id"`
		GuildName string               `bson:"name"`
		Settings  domain.GuildSettings `bson:"settings"`
	}

	var results []Guild

	if err := s.mongoRepo.Aggregate("guilds", pipeline, &results); err != nil {
		return domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	if len(results) == 0 {
		return domain.Error{
			Severity: domain.SEVERITY_WARNING,
			Message:  fmt.Sprintf("No settings found for the guild with id {%s}", guildID),
		}
	}

	config := results[0]
	settings := config.Settings
	options := settings.Welcome

	if !options.Enabled {
		return domain.Error{
			Severity: domain.SEVERITY_WARNING,
			Message:  fmt.Sprintf("Welcome plugin is not enabled for guild with id {%s}", guildID),
		}
	}

	templateKeys := map[string]string{
		"guild_name":   config.GuildName,
		"user_mention": event.Member.Mention(),
	}

	if options.Message.Type == domain.GUILD_SETTINGS_WELCOME_MESSAGE_TYPE_EMBED {
		content := options.Message.Content

		embed := &discordgo.MessageEmbed{
			Title:       utils.ParseTemplate(content.Title, templateKeys),
			Description: utils.ParseTemplate(content.Description, templateKeys),
			Color:       content.Color,
			Image: &discordgo.MessageEmbedImage{
				URL: content.Image,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: event.AvatarURL("256"),
			},
		}
		err := s.discordRepo.SendEmbed(options.ChannelID, embed)
		if err != nil {
			return domain.Error{
				Severity: domain.SEVERITY_ERROR,
				Message:  err.Error(),
				Error:    err,
			}
		}
	}

	if options.Message.Type == domain.GUILD_SETTINGS_WELCOME_MESSAGE_TYPE_TEXT {
		content := options.Message.Content

		err := s.discordRepo.SendMessage(options.ChannelID, utils.ParseTemplate(content.Description, templateKeys))
		if err != nil {
			return domain.Error{
				Severity: domain.SEVERITY_ERROR,
				Message:  err.Error(),
				Error:    err,
			}
		}

	}

	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message:  fmt.Sprintf("Welcome message sent successfully for member with id {%s}", event.Member.User.ID),
	}
}

func (s *GreetingsServiceImpl) GoodbyeMember(guildID string, event *discordgo.GuildMemberRemove) domain.Error {
	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message:  fmt.Sprintf("Welcome message sent successfully for member with id {%s}", event.Member.User.ID),
	}
}

func NewWelcomeService(discordRepo ports.DiscordRepo, mongoRepo ports.MongoRepo) *GreetingsServiceImpl {
	return &GreetingsServiceImpl{discordRepo: discordRepo, mongoRepo: mongoRepo}
}

type SetupServiceImpl struct {
	repo ports.MongoRepo
}

func (s *SetupServiceImpl) IsGuildExists(ID string) (bool, domain.Error) {
	var result domain.Guild

	ok, err := s.repo.FindOne("guilds", bson.M{"_id": ID}, &result)
	if err != nil {
		return false, domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	return ok, domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Error:    nil,
	}
}

func (s *SetupServiceImpl) LoadGuild(guild domain.Guild) domain.Error {
	if err := s.repo.InsertOne("guilds", guild); err != nil {
		return domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message:  fmt.Sprintf("Inserted guild successfully with id {%s}", guild.ID),
	}
}

func (s *SetupServiceImpl) LoadSettings(settings domain.GuildSettings) domain.Error {
	if err := s.repo.InsertOne("settings", settings); err != nil {
		return domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message:  fmt.Sprintf("Inserted settings successfully with id {%s}", settings.ID),
	}
}

func (s *SetupServiceImpl) LoadMembers(members []domain.Member) domain.Error {
	documents := make([]interface{}, len(members))
	for i, member := range members {
		documents[i] = member
	}

	if err := s.repo.InsertMany("members", documents); err != nil {
		return domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message: fmt.Sprintf("Inserted members successfully with ids {%s}", strings.Join(func() (s []string) {
			for _, v := range members {
				s = append(s, v.ID)
			}
			return s
		}(), ",")),
	}
}

func (s *SetupServiceImpl) LoadChannels(channels []domain.Channel) domain.Error {
	channelsInterface := make([]interface{}, len(channels))
	for i, member := range channels {
		channelsInterface[i] = member
	}

	if err := s.repo.InsertMany("channels", channelsInterface); err != nil {
		return domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message: fmt.Sprintf("Inserted channels successfully with ids {%s}", strings.Join(func() (s []string) {
			for _, v := range channels {
				s = append(s, v.ID)
			}
			return s
		}(), ",")),
	}
}

func (s *SetupServiceImpl) LoadRoles(roles []domain.Role) domain.Error {
	rolesInterface := make([]interface{}, len(roles))
	for i, member := range roles {
		rolesInterface[i] = member
	}

	if err := s.repo.InsertMany("roles", rolesInterface); err != nil {
		return domain.Error{
			Severity: domain.SEVERITY_ERROR,
			Message:  err.Error(),
			Error:    err,
		}
	}

	return domain.Error{
		Severity: domain.SEVERITY_SUCCESS,
		Message: fmt.Sprintf("Inserted roles successfully with ids {%s}", strings.Join(func() (s []string) {
			for _, v := range roles {
				s = append(s, v.ID)
			}
			return s
		}(), ",")),
	}
}

func NewSetupService(repo ports.MongoRepo) *SetupServiceImpl {
	return &SetupServiceImpl{repo: repo}
}
