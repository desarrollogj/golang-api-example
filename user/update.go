package user

import (
	"fmt"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/desarrollogj/golang-api-example/infrastructure"
	"github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/desarrollogj/golang-api-example/libs/logger"
)

// Update represents the method to be implemented to update an user
type Update interface {
	Execute(input domain.UserUpdateInput) (domain.User, error)
}

// defaultUpdate is the default implementation of Update interface
type defaultUpdate struct {
	repository infrastructure.UserRepository
}

// NewDefaultUpdate creates a defaultUpdate instance
func NewDefaultUpdate(repository infrastructure.UserRepository) defaultUpdate {
	return defaultUpdate{
		repository: repository,
	}
}

// Execute update an User
func (s defaultUpdate) Execute(input domain.UserUpdateInput) (domain.User, error) {
	currentUser, err := s.repository.FindByReference(input.Reference)
	if err != nil {
		errMsg := fmt.Sprintf("unexpected error when try to get user with reference %s", input.Reference)
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.NewFatalError(errMsg)
	}
	if len(currentUser.Reference) == 0 {
		return domain.User{}, errors.NewNotFoundError("user not found")
	}

	currentUser.FirstName = input.FirstName
	currentUser.LastName = input.LastName
	currentUser.Email = input.Email
	currentUser.IsActive = true
	currentUser.UpdatedDate = time.Now().UTC()

	updated, err := s.repository.Update(currentUser)
	if err != nil {
		errMsg := "unexpected error when update the user"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.NewFatalError(errMsg)
	}

	return updated, nil
}
