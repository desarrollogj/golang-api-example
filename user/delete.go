package user

import (
	"fmt"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/desarrollogj/golang-api-example/infrastructure"
	"github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/desarrollogj/golang-api-example/libs/logger"
)

// Delete represents the method to be implemented to delete (inactive) an user
type Delete interface {
	Execute(reference string) (domain.User, error)
}

// defaultDelete is the default implementation of Delete interface
type defaultDelete struct {
	repository infrastructure.UserRepository
}

// NewDefaultDelete creates a defaultDelete instance
func NewDefaultDelete(repository infrastructure.UserRepository) defaultDelete {
	return defaultDelete{
		repository: repository,
	}
}

// Execute delete an User
func (s defaultDelete) Execute(reference string) (domain.User, error) {
	currentUser, err := s.repository.FindActiveByReference(reference)
	if err != nil {
		errMsg := fmt.Sprintf("unexpected error when try to get user with reference %s", reference)
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.NewFatalError(errMsg)
	}
	if len(currentUser.Reference) == 0 {
		return domain.User{}, errors.NewNotFoundError("user not found")
	}

	currentUser.IsActive = false
	currentUser.UpdatedDate = time.Now().UTC()

	deleted, err := s.repository.Update(currentUser)
	if err != nil {
		errMsg := "unexpected error when delete the user"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.NewFatalError(errMsg)
	}

	return deleted, nil
}
