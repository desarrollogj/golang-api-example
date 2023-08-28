package user

import (
	"errors"
	"testing"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDelete_GivenAReference_WhenExecute_ThenDeleteAnUser(t *testing.T) {
	t.Log("Successfully delete an User")

	reference := "REF1"
	currentUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference: reference,
			IsActive:  true,
		},
	}
	deletedUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference: reference,
			IsActive:  false,
		},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindActiveByReference", reference).Return(currentUser, nil)
	repositoryMock.On("Update", mock.AnythingOfType("User")).Return(deletedUser, nil)

	useCase := NewDefaultDelete(repositoryMock)

	deleted, err := useCase.Execute(reference)

	assert.Nil(t, err)
	assert.NotNil(t, deleted)
	assert.Equal(t, deletedUser, deleted)

	repositoryMock.AssertExpectations(t)
}

func TestDelete_GivenAReference_WhenExecuteAndFindReturnedAnError_ThenReturnAnError(t *testing.T) {
	t.Log("Failure to delete an User because find returned an unexpected error")

	reference := "REF1"
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindActiveByReference", reference).Return(domain.User{}, errors.New("repository error"))

	useCase := NewDefaultDelete(repositoryMock)

	_, err := useCase.Execute(reference)

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected error when try to get user with reference REF1", err.Error())

	repositoryMock.AssertExpectations(t)
}

func TestDelete_GivenAReference_WhenExecuteAndUserNotFound_ThenReturnAnError(t *testing.T) {
	t.Log("Failure to delete an User because user to delete uwas not found")

	reference := "REF1"
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindActiveByReference", reference).Return(domain.User{}, nil)

	useCase := NewDefaultDelete(repositoryMock)

	_, err := useCase.Execute(reference)

	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())

	repositoryMock.AssertExpectations(t)
}

func TestDelete_GivenAReference_WhenExecuteAndUpdateReturnedAnError_ThenReturnAFatalError(t *testing.T) {
	t.Log("Failure to delete an User because update returned an unexpected error")

	reference := "REF1"
	currentUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference: reference,
			IsActive:  true,
		},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindActiveByReference", reference).Return(currentUser, nil)
	repositoryMock.On("Update", mock.AnythingOfType("User")).Return(domain.User{}, errors.New("repository error"))

	useCase := NewDefaultDelete(repositoryMock)

	_, err := useCase.Execute(reference)

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected error when delete the user", err.Error())

	repositoryMock.AssertExpectations(t)
}
