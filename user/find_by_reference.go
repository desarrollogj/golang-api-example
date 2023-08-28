package user

import (
	"fmt"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/desarrollogj/golang-api-example/infrastructure"
	"github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/desarrollogj/golang-api-example/libs/logger"
)

// FindByReference represents the method to be implemented to get an user by its reference
type FindByReference interface {
	Execute(reference string) (domain.User, error)
}

// defaultFindByReference is the default implementation of FindByReference interface
type defaultFindByReference struct {
	repository infrastructure.UserRepository
}

// NewDefaultFindByReference creates a DefaultGetById instance
func NewDefaultFindByReference(repository infrastructure.UserRepository) defaultFindByReference {
	return defaultFindByReference{
		repository: repository,
	}
}

// Execute get an user by its reference
func (s defaultFindByReference) Execute(reference string) (domain.User, error) {
	user, err := s.repository.FindActiveByReference(reference)
	if err != nil {
		errMsg := fmt.Sprintf("unexpected error when try to get user with reference %s", reference)
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.NewFatalError(errMsg)
	}

	if len(user.Reference) == 0 {
		return domain.User{}, errors.NewNotFoundError("user not found")
	}

	return user, nil
}
