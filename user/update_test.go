package user

import (
	"errors"
	"testing"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate_GivenAnUser_WhenExecute_ThenUpdateAnUser(t *testing.T) {
	t.Log("Successfully update an User")

	reference := "REF1"
	input := domain.UserUpdateInput{
		UserCreateInput: domain.UserCreateInput{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
		Reference: reference,
	}
	currentUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference: reference,
		},
		FirstName: "Another Foo",
		LastName:  "Another Bar",
		Email:     "anotherfoobar@email.com",
	}
	updatedUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference: reference,
		},
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindByReference", reference).Return(currentUser, nil)
	repositoryMock.On("Update", mock.AnythingOfType("User")).Return(updatedUser, nil)

	useCase := NewDefaultUpdate(repositoryMock)

	updated, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, updatedUser, updated)

	repositoryMock.AssertExpectations(t)
}

func TestUpdate_GivenAnUser_WhenExecuteAndFindReturnedAnError_ThenReturnAnError(t *testing.T) {
	t.Log("Failure to update an User because find returned an unexpected error")

	reference := "REF1"
	input := domain.UserUpdateInput{
		UserCreateInput: domain.UserCreateInput{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
		Reference: reference,
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindByReference", reference).Return(domain.User{}, errors.New("repository error"))

	useCase := NewDefaultUpdate(repositoryMock)

	_, err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected error when try to get user with reference REF1", err.Error())

	repositoryMock.AssertExpectations(t)
}

func TestUpdate_GivenAnUser_WhenExecuteAndUserNotFound_ThenReturnANotFoundError(t *testing.T) {
	t.Log("Failure to update an User because user to update was not found")

	reference := "REF1"
	input := domain.UserUpdateInput{
		UserCreateInput: domain.UserCreateInput{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
		Reference: reference,
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindByReference", reference).Return(domain.User{}, nil)

	useCase := NewDefaultUpdate(repositoryMock)

	_, err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())

	repositoryMock.AssertExpectations(t)
}

func TestUpdate_GivenAnUser_WhenExecuteAndUpdateReturnedAnError_ThenReturnAFatalError(t *testing.T) {
	t.Log("Failure to update an User because update returned an unexpected error")

	reference := "REF1"
	input := domain.UserUpdateInput{
		UserCreateInput: domain.UserCreateInput{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
		Reference: reference,
	}
	currentUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference: reference,
		},
		FirstName: "Another Foo",
		LastName:  "Another Bar",
		Email:     "anotherfoobar@email.com",
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindByReference", reference).Return(currentUser, nil)
	repositoryMock.On("Update", mock.AnythingOfType("User")).Return(domain.User{}, errors.New("repository error"))

	useCase := NewDefaultUpdate(repositoryMock)

	_, err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected error when update the user", err.Error())

	repositoryMock.AssertExpectations(t)
}
