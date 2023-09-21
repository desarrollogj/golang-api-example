package user

import (
	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/desarrollogj/golang-api-example/infrastructure"
	"github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/desarrollogj/golang-api-example/libs/logger"
)

// Search represents the method to be implemented to search users
type Search interface {
	Execute(input domain.UserSearchInput) (domain.UserSearchOutput, error)
}

// defaultSearch is the default implementation of Search interface
type defaultSearch struct {
	repository infrastructure.UserRepository
}

// NewDefaulSearch creates a defaultSearch instance
func NewDefaulSearch(repository infrastructure.UserRepository) defaultSearch {
	return defaultSearch{
		repository: repository,
	}
}

// Execute Search users
func (s defaultSearch) Execute(input domain.UserSearchInput) (domain.UserSearchOutput, error) {
	output, err := s.repository.SearchActive(input)
	if err != nil {
		errMsg := "unexpected error when try to search users"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.UserSearchOutput{}, errors.NewFatalError(errMsg)
	}
	return output, nil
}
