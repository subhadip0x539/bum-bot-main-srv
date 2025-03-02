package services

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/bwmarrin/discordgo"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/domain"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/ports"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/utils"
)

type GreetingsServiceImpl struct {
	discordRepo ports.DiscordRepo
	mongoRepo   ports.MongoRepo
}

func (s *GreetingsServiceImpl) AddMember(member domain.Member) *domain.Error {
	if err := s.mongoRepo.InsertOne("members", member); err != nil {
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}
	return nil
}

func (s *GreetingsServiceImpl) RemoveMember(memberID string, guildID string) *domain.Error {
	if err := s.mongoRepo.DeleteOne("members", bson.M{"user_id": memberID, "guild_id": guildID}); err != nil {
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}
	return nil
}

func (s *GreetingsServiceImpl) WelcomeMember(guildID string, event *discordgo.GuildMemberAdd) *domain.Error {
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
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}

	if len(results) == 0 {
		err := fmt.Errorf("no settings found for the guild with id {%s}", guildID)
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}

	config := results[0]
	settings := config.Settings
	plugin := settings.Plugins[0]

	if !plugin.Enabled {
		return domain.NewError(nil, fmt.Sprintf("Welcome plugin is not enabled for guild with id {%s}", guildID), domain.SEVERITY_SUCCESS)
	}

	options, ok := plugin.Options.(domain.GuildSettingsPluginWelcomeOptions)
	if !ok {
		err := fmt.Errorf("invalid welcome plugin options for guild with id {%s}", guildID)
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
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
			return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
		}
	}

	if options.Message.Type == domain.GUILD_SETTINGS_WELCOME_MESSAGE_TYPE_TEXT {
		content := options.Message.Content

		err := s.discordRepo.SendMessage(options.ChannelID, utils.ParseTemplate(content.Description, templateKeys))
		if err != nil {
			return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
		}
	}

	return nil
}

func (s *GreetingsServiceImpl) GoodbyeMember(guildID string, event *discordgo.GuildMemberRemove) *domain.Error {
	return nil
}

func NewWelcomeService(discordRepo ports.DiscordRepo, mongoRepo ports.MongoRepo) *GreetingsServiceImpl {
	return &GreetingsServiceImpl{discordRepo: discordRepo, mongoRepo: mongoRepo}
}

type SetupServiceImpl struct {
	repo ports.MongoRepo
}

func (s *SetupServiceImpl) IsGuildExists(ID string) (bool, *domain.Error) {
	var result domain.Guild

	ok, err := s.repo.FindOne("guilds", bson.M{"_id": ID}, &result)
	if err != nil {
		return false, domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}

	return ok, nil
}

func (s *SetupServiceImpl) LoadGuild(guild domain.Guild) *domain.Error {
	if err := s.repo.InsertOne("guilds", guild); err != nil {
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}
	return nil
}

func (s *SetupServiceImpl) LoadSettings(settings domain.GuildSettings) *domain.Error {
	if err := s.repo.InsertOne("settings", settings); err != nil {
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}
	return nil
}

func (s *SetupServiceImpl) LoadMembers(members []domain.Member) *domain.Error {
	documents := make([]interface{}, len(members))
	for i, member := range members {
		documents[i] = member
	}

	if err := s.repo.InsertMany("members", documents); err != nil {
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}

	return nil
}

func (s *SetupServiceImpl) LoadChannels(channels []domain.Channel) *domain.Error {
	channelsInterface := make([]interface{}, len(channels))
	for i, member := range channels {
		channelsInterface[i] = member
	}

	if err := s.repo.InsertMany("channels", channelsInterface); err != nil {
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}

	return nil
}

func (s *SetupServiceImpl) LoadRoles(roles []domain.Role) *domain.Error {
	rolesInterface := make([]interface{}, len(roles))
	for i, member := range roles {
		rolesInterface[i] = member
	}

	if err := s.repo.InsertMany("roles", rolesInterface); err != nil {
		return domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}

	return nil
}

func (s *SetupServiceImpl) GetPlugins() ([]domain.Plugin, *domain.Error) {
	var result []domain.Plugin

	if err := s.repo.FindAll("plugins", bson.M{}, &result); err != nil {
		return []domain.Plugin{{}}, domain.NewError(err, err.Error(), domain.SEVERITY_ERROR)
	}

	return result, nil
}

func NewSetupService(repo ports.MongoRepo) *SetupServiceImpl {
	return &SetupServiceImpl{repo: repo}
}
