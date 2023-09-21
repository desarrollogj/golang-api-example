package handler

import (
	"errors"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/mock"
)

// Config
func newApplicationConfigurationMock() domain.ApplicationConfiguration {
	return domain.ApplicationConfiguration{
		PagingDefaultPage: 1,
		PagingDefaultSize: 10,
	}
}

// Mapper
type userMapperMock struct {
	mock.Mock
}

func (m *userMapperMock) MapDomainToResponse(user domain.User) UserResponse {
	args := m.Called(user)

	t, ok := args.Get(0).(UserResponse)
	if !ok {
		return UserResponse{}
	}

	return t
}

func (m *userMapperMock) MapDomainListToResponseList(users []domain.User) []UserResponse {
	args := m.Called(users)

	t, ok := args.Get(0).([]UserResponse)
	if !ok {
		return []UserResponse{}
	}

	return t
}

func (m *userMapperMock) MapCreateRequestToInput(request UserCreateRequest) domain.UserCreateInput {
	args := m.Called(request)

	t, ok := args.Get(0).(domain.UserCreateInput)
	if !ok {
		return domain.UserCreateInput{}
	}

	return t
}

func (m *userMapperMock) MapUpdateRequestToInput(reference string, request UserUpdateRequest) domain.UserUpdateInput {
	args := m.Called(reference, request)

	t, ok := args.Get(0).(domain.UserUpdateInput)
	if !ok {
		return domain.UserUpdateInput{}
	}

	return t
}

func (m *userMapperMock) MapDomainSearchOutputToResponse(output domain.UserSearchOutput) UserSearchResponse {
	args := m.Called(output)

	t, ok := args.Get(0).(UserSearchResponse)
	if !ok {
		return UserSearchResponse{}
	}

	return t
}

// Services
type userCreateServiceMock struct {
	mock.Mock
}

func (s *userCreateServiceMock) Execute(input domain.UserCreateInput) (domain.User, error) {
	args := s.Called(input)

	t, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock_error")
	}

	return t, args.Error(1)
}

type userDeleteServiceMock struct {
	mock.Mock
}

func (s *userDeleteServiceMock) Execute(reference string) (domain.User, error) {
	args := s.Called(reference)

	t, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock_error")
	}

	return t, args.Error(1)
}

type userFindAllServiceMock struct {
	mock.Mock
}

func (s *userFindAllServiceMock) Execute() ([]domain.User, error) {
	args := s.Called()

	t, ok := args.Get(0).([]domain.User)
	if !ok {
		return []domain.User{}, errors.New("mock_error")
	}

	return t, args.Error(1)
}

type userFindByReferenceServiceMock struct {
	mock.Mock
}

func (s *userFindByReferenceServiceMock) Execute(reference string) (domain.User, error) {
	args := s.Called(reference)

	t, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock_error")
	}

	return t, args.Error(1)
}

type userUpdateServiceMock struct {
	mock.Mock
}

func (s *userUpdateServiceMock) Execute(input domain.UserUpdateInput) (domain.User, error) {
	args := s.Called(input)

	t, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock_error")
	}

	return t, args.Error(1)
}

type userSearchServiceMock struct {
	mock.Mock
}

func (s *userSearchServiceMock) Execute(input domain.UserSearchInput) (domain.UserSearchOutput, error) {
	args := s.Called(input)

	t, ok := args.Get(0).(domain.UserSearchOutput)
	if !ok {
		return domain.UserSearchOutput{}, errors.New("mock_error")
	}

	return t, args.Error(1)
}
