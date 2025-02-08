package services

import (
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/domain"
	ports "github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/ports"
)

type SetupServiceImpl struct {
	repo ports.MongoRepo
}

func (s *SetupServiceImpl) LoadServer(guild domain.Guild) error {
	if err := s.repo.InsertOne("guilds", guild); err != nil {
		return err
	}
	return nil
}

func (s *SetupServiceImpl) LoadMembers(members []domain.Member) error {
	membersInterface := make([]interface{}, len(members))
	for i, member := range members {
		membersInterface[i] = member
	}

	if err := s.repo.InsertMany("members", membersInterface); err != nil {
		return err
	}
	return nil
}

func (s *SetupServiceImpl) LoadChannels(channels []domain.Channel) error {
	channelsInterface := make([]interface{}, len(channels))
	for i, member := range channels {
		channelsInterface[i] = member
	}

	if err := s.repo.InsertMany("channels", channelsInterface); err != nil {
		return err
	}
	return nil
}

func (s *SetupServiceImpl) LoadRoles(roles []domain.Role) error {
	rolesInterface := make([]interface{}, len(roles))
	for i, member := range roles {
		rolesInterface[i] = member
	}

	if err := s.repo.InsertMany("roles", rolesInterface); err != nil {
		return err
	}
	return nil
}

func NewSetupService(repo ports.MongoRepo) *SetupServiceImpl {
	return &SetupServiceImpl{repo: repo}
}
