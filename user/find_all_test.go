package user

import (
	"errors"
	"testing"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/assert"
)

func TestFindAll_WhenExecute_ThenGetAnActiveUserList(t *testing.T) {
	t.Log("Successfully find all Users")

	users := []domain.User{
		{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindAllActive").Return(users, nil)

	useCase := NewDefaultFindAll(repositoryMock)

	foundUsers, err := useCase.Execute()

	assert.Nil(t, err)
	assert.NotNil(t, foundUsers)
	assert.Equal(t, users, foundUsers)

	repositoryMock.AssertExpectations(t)
}

func TestFindAll_WhenExecute_AndRepositoryReturnedAnError_ThenReturnAFatalError(t *testing.T) {
	t.Log("Failure to find all Users because repository returned an error")

	repositoryMock := new(repositoryMock)
	repositoryMock.On("FindAllActive").Return([]domain.User{}, errors.New("repository error"))

	useCase := NewDefaultFindAll(repositoryMock)

	_, err := useCase.Execute()

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected error when find all users", err.Error())

	repositoryMock.AssertExpectations(t)
}
