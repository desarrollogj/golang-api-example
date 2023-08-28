package infrastructure

import (
	"testing"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/assert"
)

func TestMongoUserRepositoryMapper_GivenDomainData_WhenMap_ThenMapToRepositoryData(t *testing.T) {
	t.Log("Should map user domain data to user repository data")

	now := time.Now().UTC()
	domainUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference:   "USER1",
			IsActive:    true,
			CreatedDate: now,
			UpdatedDate: now,
		},
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@test.com",
	}
	expectedRepoUser := MongoUser{
		Reference:   "USER1",
		FirstName:   "Foo",
		LastName:    "Bar",
		Email:       "foobar@test.com",
		IsActive:    true,
		CreatedDate: now,
		UpdatedDate: now,
	}

	mapper := NewDefaultMongoRepositoryMapper()
	repoUser := mapper.MapDomainToRepository(domainUser)

	assert.NotNil(t, repoUser)
	assert.Equal(t, expectedRepoUser, repoUser)
}

func TestMongoUserRepositoryMapper_GivenRepositoryData_WhenMap_ThenMapToDomainData(t *testing.T) {
	t.Log("Should map user repository data to user domain data")

	now := time.Now().UTC()
	repositoryUser := MongoUser{
		Reference:   "USER1",
		FirstName:   "Foo",
		LastName:    "Bar",
		Email:       "foobar@test.com",
		IsActive:    true,
		CreatedDate: now,
		UpdatedDate: now,
	}
	expectedDomainUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference:   "USER1",
			IsActive:    true,
			CreatedDate: now,
			UpdatedDate: now,
		},
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@test.com",
	}

	mapper := NewDefaultMongoRepositoryMapper()
	domainUser := mapper.MapRepositoryToDomain(repositoryUser)

	assert.NotNil(t, domainUser)
	assert.Equal(t, expectedDomainUser, domainUser)
}

func TestMongoUserRepositoryMapper_GivenRepositoryDataList_WhenMap_ThenMapToDomainDataList(t *testing.T) {
	t.Log("Should map user repository data list to user domain data list")

	now := time.Now().UTC()
	repositoryUsers := []MongoUser{
		{
			Reference:   "USER1",
			FirstName:   "Foo",
			LastName:    "Bar",
			Email:       "foobar@test.com",
			IsActive:    true,
			CreatedDate: now,
			UpdatedDate: now,
		},
	}
	expectedDomainUsers := []domain.User{
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
	}

	mapper := NewDefaultMongoRepositoryMapper()
	domainUser := mapper.MapRepositoryListToDomainList(repositoryUsers)

	assert.NotNil(t, domainUser)
	assert.Equal(t, expectedDomainUsers, domainUser)
}
