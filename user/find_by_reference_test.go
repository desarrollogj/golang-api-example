package user

import (
	"errors"
	"testing"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/assert"
)

func TestFindByReference_GivenAReference_WhenExecute_ThenGetAnUser(t *testing.T) {
	t.Log("Successfully get an User by its reference")

	reference := "USER1"
	user := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference: reference,
			IsActive:  true,
		},
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindActiveByReference", reference).Return(user, nil)

	useCase := NewDefaultFindByReference(repositoryMock)

	foundUser, err := useCase.Execute(reference)

	assert.Nil(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user, foundUser)

	repositoryMock.AssertExpectations(t)
}

func TestFindByReference_GivenAReference_WhenExecute_AndUsesDoesNotExist_ThenReturnNotFoundError(t *testing.T) {
	t.Log("Failure to get an User by its reference")

	reference := "USER1"
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindActiveByReference", reference).Return(domain.User{}, nil)

	useCase := NewDefaultFindByReference(repositoryMock)

	_, err := useCase.Execute(reference)

	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())

	repositoryMock.AssertExpectations(t)
}

func TestFindByReference_GivenAReference_WhenExecute_AndRepositoryRetunedAnError_ThenReturnFatalError(t *testing.T) {
	t.Log("Failure to get an User by its reference because repository returned an error")

	reference := "USER1"
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindActiveByReference", reference).Return(domain.User{}, errors.New("repository error"))

	useCase := NewDefaultFindByReference(repositoryMock)

	_, err := useCase.Execute(reference)

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected error when try to get user with reference USER1", err.Error())

	repositoryMock.AssertExpectations(t)
}
