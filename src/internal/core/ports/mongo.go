package ports

import "github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/domain"

type MongoRepo interface {
	InsertOne(collation string, document interface{}) error
	InsertMany(collation string, document []interface{}) error
}

type SetupService interface {
	LoadServer(server domain.Guild) error
	LoadMembers(members []domain.Member) error
	LoadChannels(channels []domain.Channel) error
	LoadRoles(roles []domain.Role) error
}
