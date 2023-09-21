package handler

import (
	"testing"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserMapper_GivenAUserDomain_WhenMapDomainToResponse_ThenReturnUserResponse(t *testing.T) {
	t.Log("Successfully map domain user to response user")

	current := time.Now().UTC()
	currentStr := current.Format(time.RFC3339)
	domainUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference:   "USER1",
			IsActive:    true,
			CreatedDate: current,
			UpdatedDate: current,
		},
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	responseUser := UserResponse{
		Id:          "USER1",
		FirstName:   "Foo",
		LastName:    "Bar",
		Email:       "foobar@email.com",
		IsActive:    true,
		CreatedDate: currentStr,
		UpdatedDate: currentStr,
	}

	mapper := NewDefaultUserMapper()
	response := mapper.MapDomainToResponse(domainUser)

	assert.NotNil(t, response)
	assert.Equal(t, responseUser, response)
}

func TestUserMapper_GivenAUserDomainList_WhenMapDomainListToResponse_ThenReturnUserResponseList(t *testing.T) {
	t.Log("Successfully map domain user list to response user list")

	current := time.Now().UTC()
	currentStr := current.Format(time.RFC3339)
	domainUsers := []domain.User{
		{
			GenericEntity: domain.GenericEntity{
				Reference:   "USER1",
				IsActive:    true,
				CreatedDate: current,
				UpdatedDate: current,
			},
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
	}
	responseUsers := []UserResponse{
		{
			Id:          "USER1",
			FirstName:   "Foo",
			LastName:    "Bar",
			Email:       "foobar@email.com",
			IsActive:    true,
			CreatedDate: currentStr,
			UpdatedDate: currentStr,
		},
	}

	mapper := NewDefaultUserMapper()
	response := mapper.MapDomainListToResponseList(domainUsers)

	assert.NotNil(t, response)
	assert.Equal(t, responseUsers, response)
}

func TestUserMapper_GivenACreateUserRequest_WhenMapRequestToDomain_ThenReturnCreateUserDomain(t *testing.T) {
	t.Log("Successfully map create user request to input")

	request := UserCreateRequest{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	expectedInput := domain.UserCreateInput{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}

	mapper := NewDefaultUserMapper()
	input := mapper.MapCreateRequestToInput(request)

	assert.NotNil(t, input)
	assert.Equal(t, expectedInput, input)
}

func TestUserMapper_GivenAnUpdateUserRequest_WhenMapRequestToDomain_ThenReturnUpdateUserDomain(t *testing.T) {
	t.Log("Successfully map update user request to input")

	reference := "USER1"
	request := UserUpdateRequest{
		UserCreateRequest: UserCreateRequest{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
	}
	expectedInput := domain.UserUpdateInput{
		UserCreateInput: domain.UserCreateInput{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
		Reference: reference,
	}

	mapper := NewDefaultUserMapper()
	input := mapper.MapUpdateRequestToInput(reference, request)

	assert.NotNil(t, input)
	assert.Equal(t, expectedInput, input)
}

func TestUserMapper_GivenASearchOutputDomain_WhenMapDomainToResponse_ThenReturnSearchOutputResponse(t *testing.T) {
	t.Log("Successfully map domain user list to response user list")

	total := int64(1)
	page := 1
	size := 10
	current := time.Now().UTC()
	currentStr := current.Format(time.RFC3339)
	domainSearchOutput := domain.UserSearchOutput{
		SearchOutput: domain.SearchOutput{
			Total:    total,
			Page:     page,
			PageSize: size,
		},
		Users: []domain.User{
			{
				GenericEntity: domain.GenericEntity{
					Reference:   "USER1",
					IsActive:    true,
					CreatedDate: current,
					UpdatedDate: current,
				},
				FirstName: "Foo",
				LastName:  "Bar",
				Email:     "foobar@email.com",
			},
		},
	}
	searchResponse := UserSearchResponse{
		Data: []UserResponse{
			{
				Id:          "USER1",
				FirstName:   "Foo",
				LastName:    "Bar",
				Email:       "foobar@email.com",
				IsActive:    true,
				CreatedDate: currentStr,
				UpdatedDate: currentStr,
			},
		},
		Total:    total,
		Page:     page,
		PageSize: size,
	}

	mapper := NewDefaultUserMapper()
	response := mapper.MapDomainSearchOutputToResponse(domainSearchOutput)

	assert.NotNil(t, response)
	assert.Equal(t, searchResponse, response)
}
