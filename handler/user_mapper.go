package handler

import (
	"strings"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
)

type UserResponse struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	IsActive    bool   `json:"isActive"`
	CreatedDate string `json:"created"`
	UpdatedDate string `json:"updated"`
}

type UserCreateRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

type UserUpdateRequest struct {
	UserCreateRequest
}

// UserMapper represents the method for user mappers
type UserMapper interface {
	MapDomainToResponse(user domain.User) UserResponse
	MapDomainListToResponseList(users []domain.User) []UserResponse
	MapCreateRequestToInput(request UserCreateRequest) domain.UserCreateInput
	MapUpdateRequestToInput(reference string, request UserUpdateRequest) domain.UserUpdateInput
}

// defaultUserMapper is the default implementation for UserMapper interface
type defaultUserMapper struct {
}

// NewDefaultUserMapper creates a defaultUserMapper handler
func NewDefaultUserMapper() defaultUserMapper {
	return defaultUserMapper{}
}

// MapDomainToResponse mas a domain user to a response
func (m defaultUserMapper) MapDomainToResponse(user domain.User) UserResponse {
	return UserResponse{
		Id:          user.Reference,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsActive:    user.IsActive,
		CreatedDate: user.CreatedDate.UTC().Format(time.RFC3339),
		UpdatedDate: user.UpdatedDate.UTC().Format(time.RFC3339),
	}
}

// MapDomainListToResponseList map a list of domain users to a response list
func (m defaultUserMapper) MapDomainListToResponseList(users []domain.User) []UserResponse {
	usersResponse := []UserResponse{}

	for _, user := range users {
		usersResponse = append(usersResponse, m.MapDomainToResponse(user))
	}

	return usersResponse
}

// MapCreateRequestToInput map create request to an input struct
func (m defaultUserMapper) MapCreateRequestToInput(request UserCreateRequest) domain.UserCreateInput {
	return domain.UserCreateInput{
		FirstName: strings.TrimSpace(request.FirstName),
		LastName:  strings.TrimSpace(request.LastName),
		Email:     strings.TrimSpace(request.Email),
	}
}

// MapUpdateRequestToInput map update request to an input struct
func (m defaultUserMapper) MapUpdateRequestToInput(reference string, request UserUpdateRequest) domain.UserUpdateInput {
	return domain.UserUpdateInput{
		UserCreateInput: domain.UserCreateInput{
			FirstName: strings.TrimSpace(request.FirstName),
			LastName:  strings.TrimSpace(request.LastName),
			Email:     strings.TrimSpace(request.Email),
		},
		Reference: strings.TrimSpace(reference),
	}
}
