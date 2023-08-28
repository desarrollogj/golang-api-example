package user

import (
	"errors"

	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) FindAllActive() ([]domain.User, error) {
	args := m.Called()

	users, ok := args.Get(0).([]domain.User)
	if !ok {
		return []domain.User{}, errors.New("mock error")
	}

	return users, args.Error(1)
}

func (m *repositoryMock) FindActiveByReference(reference string) (domain.User, error) {
	args := m.Called(reference)

	user, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock error")
	}

	return user, args.Error(1)
}

func (m *repositoryMock) FindByReference(reference string) (domain.User, error) {
	args := m.Called(reference)

	user, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock error")
	}

	return user, args.Error(1)
}

func (m *repositoryMock) Create(user domain.User) (domain.User, error) {
	args := m.Called(user)

	user, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock error")
	}

	return user, args.Error(1)
}

func (m *repositoryMock) Update(user domain.User) (domain.User, error) {
	args := m.Called(user)

	user, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock error")
	}

	return user, args.Error(1)
}

func (m *repositoryMock) Delete(reference string) (domain.User, error) {
	args := m.Called(reference)

	user, ok := args.Get(0).(domain.User)
	if !ok {
		return domain.User{}, errors.New("mock error")
	}

	return user, args.Error(1)
}
