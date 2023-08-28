package user

import (
	"errors"
	"testing"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate_GivenAnUser_WhenExecute_ThenCreateAnUser(t *testing.T) {
	t.Log("Successfully create a User")

	input := domain.UserCreateInput{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	createdUser := domain.User{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Create", mock.AnythingOfType("User")).Return(createdUser, nil)

	useCase := NewDefaultCreate(repositoryMock)

	created, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, createdUser, created)

	repositoryMock.AssertExpectations(t)
}

func TestCreate_GivenAnUser_WhenExecute_AndRepositoryFailred_ThenReturnAFatalError(t *testing.T) {
	t.Log("Failure to create an User because repository returned an error")

	input := domain.UserCreateInput{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Create", mock.AnythingOfType("User")).Return(domain.User{}, errors.New("repository error"))

	useCase := NewDefaultCreate(repositoryMock)

	_, err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected error when create the user", err.Error())

	repositoryMock.AssertExpectations(t)
}
