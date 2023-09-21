package infrastructure

import "github.com/desarrollogj/golang-api-example/domain"

// UserMongoRepositoryMapper represents the methods to be implemented by mongo domain entities mapper
type UserMongoRepositoryMapper interface {
	MapDomainToRepository(user domain.User) MongoUser
	MapRepositoryToDomain(user MongoUser) domain.User
	MapRepositoryListToDomainList(users []MongoUser) []domain.User
	MapRepositorySearchActiveToOutput(users []MongoUser, total int64, page int, size int) domain.UserSearchOutput
}

// defaultMongoRepositoryMapper is the default implementation of UserMongoRepositoryMapper
type defaultMongoRepositoryMapper struct {
}

// NewDefaultMongoRepositoryMapper creates a new defaultMongoRepositoryMapper
func NewDefaultMongoRepositoryMapper() defaultMongoRepositoryMapper {
	return defaultMongoRepositoryMapper{}
}

func (m defaultMongoRepositoryMapper) MapDomainToRepository(user domain.User) MongoUser {
	return MongoUser{
		Reference:   user.Reference,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsActive:    user.IsActive,
		CreatedDate: user.CreatedDate,
		UpdatedDate: user.UpdatedDate,
	}
}

func (m defaultMongoRepositoryMapper) MapRepositoryToDomain(user MongoUser) domain.User {
	return domain.User{
		GenericEntity: domain.GenericEntity{
			Reference:   user.Reference,
			IsActive:    user.IsActive,
			CreatedDate: user.CreatedDate,
			UpdatedDate: user.UpdatedDate,
		},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func (m defaultMongoRepositoryMapper) MapRepositoryListToDomainList(users []MongoUser) []domain.User {
	mappedUsers := []domain.User{}

	for _, user := range users {
		mappedUsers = append(mappedUsers, m.MapRepositoryToDomain(user))
	}

	return mappedUsers
}

func (m defaultMongoRepositoryMapper) MapRepositorySearchActiveToOutput(users []MongoUser, total int64, page int, size int) domain.UserSearchOutput {
	return domain.UserSearchOutput{
		SearchOutput: domain.SearchOutput{
			Total:    total,
			Page:     page,
			PageSize: size,
		},
		Users: m.MapRepositoryListToDomainList(users),
	}
}
