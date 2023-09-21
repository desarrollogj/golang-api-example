package user

import (
	"errors"
	"testing"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/assert"
)

func TestSearch_WhenExecute_ThenGetSearchResultOutput(t *testing.T) {
	t.Log("Successfully search Users")

	now := time.Now().UTC()
	searchInput := domain.UserSearchInput{
		SearchInput: domain.SearchInput{
			Page:     1,
			PageSize: 10,
		},
	}
	searchOutput := domain.UserSearchOutput{
		SearchOutput: domain.SearchOutput{
			Total:    1,
			Page:     1,
			PageSize: 10,
		},
		Users: []domain.User{
			{
				GenericEntity: domain.GenericEntity{
					Reference:   "USER1",
					IsActive:    true,
					CreatedDate: now,
					UpdatedDate: now,
				},
				FirstName: "Foo",
				LastName:  "Bar",
				Email:     "foobar@test.com",
			},
		},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("SearchActive", searchInput).Return(searchOutput, nil)

	useCase := NewDefaulSearch(repositoryMock)

	usersFound, err := useCase.Execute(searchInput)

	assert.Nil(t, err)
	assert.NotNil(t, usersFound)
	assert.Equal(t, searchOutput, usersFound)

	repositoryMock.AssertExpectations(t)
}

func TestSearch_WhenExecute_AndRepositoryReturnsError_ThenReturnFatalError(t *testing.T) {
	t.Log("Failure to search Users because repository returned an error")

	searchInput := domain.UserSearchInput{
		SearchInput: domain.SearchInput{
			Page:     1,
			PageSize: 10,
		},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("SearchActive", searchInput).Return(domain.UserSearchOutput{}, errors.New("repository error"))

	useCase := NewDefaulSearch(repositoryMock)

	_, err := useCase.Execute(searchInput)

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected error when try to search users", err.Error())

	repositoryMock.AssertExpectations(t)
}
