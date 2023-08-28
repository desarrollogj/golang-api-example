package user

import (
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/desarrollogj/golang-api-example/infrastructure"
	"github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/desarrollogj/golang-api-example/libs/logger"
	"github.com/google/uuid"
)

// Create represents the method to be implemented to create an user
type Create interface {
	Execute(input domain.UserCreateInput) (domain.User, error)
}

// defaultCreate is the default implementation of Create interface
type defaultCreate struct {
	repository infrastructure.UserRepository
}

// NewDefaultCreate creates a defaultCreate instance
func NewDefaultCreate(repository infrastructure.UserRepository) defaultCreate {
	return defaultCreate{
		repository: repository,
	}
}

// Create an User
func (s defaultCreate) Execute(input domain.UserCreateInput) (domain.User, error) {
	created := time.Now().UTC()
	user := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference:   uuid.NewString(),
			IsActive:    true,
			CreatedDate: created,
			UpdatedDate: created,
		},
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	user, err := s.repository.Create(user)
	if err != nil {
		errMsg := "unexpected error when create the user"
		logger.AppLog.Error().Err(err).Msg(errMsg)
		return domain.User{}, errors.NewFatalError(errMsg)
	}

	return user, nil
}
