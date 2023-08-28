package user

import (
	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/desarrollogj/golang-api-example/infrastructure"
	"github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/desarrollogj/golang-api-example/libs/logger"
)

// FindAll represents the method to be implemented to get all users
type FindAll interface {
	Execute() ([]domain.User, error)
}

// defaultFindAll is the default implementation of FindAll interface
type defaultFindAll struct {
	repository infrastructure.UserRepository
}

// NewDefaultFindAll creates a defaultFindAll instance
func NewDefaultFindAll(repository infrastructure.UserRepository) defaultFindAll {
	return defaultFindAll{
		repository: repository,
	}
}

// Execute Get all users
func (s defaultFindAll) Execute() ([]domain.User, error) {
	users, err := s.repository.FindAllActive()
	if err != nil {
		errMsg := "unexpected error when find all users"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return []domain.User{}, errors.NewFatalError(errMsg)
	}
	return users, nil
}
