package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	libErrors "github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUser_WhenFindAll_ThenReturnUserListResponse(t *testing.T) {
	t.Log("Successfully find all users")

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

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	mapperMock.On("MapDomainListToResponseList", domainUsers).Return(responseUsers)
	findAllMock := new(userFindAllServiceMock)
	findAllMock.On("Execute").Return(domainUsers, nil)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)

	r := testRouter()
	r.GET("/api/v1/users", handler.FindAll)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result []UserResponse
	json.NewDecoder(w.Body).Decode(&result)

	assert.NotNil(t, result)
	assert.Equal(t, responseUsers, result)

	mapperMock.AssertExpectations(t)
	findAllMock.AssertExpectations(t)
}

func TestUser_WhenFindAll_AndServiceReturnedAnError_ThenReturnInternalServerErrorResponse(t *testing.T) {
	t.Log("Failure when find all users because service returned an unexpected error")

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findAllMock.On("Execute").Return([]domain.User{}, errors.New("service error"))
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)

	r := testRouter()
	r.GET("/api/v1/users", handler.FindAll)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "service error", err.Message)

	findAllMock.AssertExpectations(t)
}

func TestUser_GivenAnId_WhenFindById_ThenReturnUserResponse(t *testing.T) {
	t.Log("Successfully find an user by its id")

	current := time.Now().UTC()
	currentStr := current.Format(time.RFC3339)
	reference := "USER1"
	domainUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference:   reference,
			IsActive:    true,
			CreatedDate: current,
			UpdatedDate: current,
		},
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	responseUser := UserResponse{
		Id:          reference,
		FirstName:   "Foo",
		LastName:    "Bar",
		Email:       "foobar@email.com",
		IsActive:    true,
		CreatedDate: currentStr,
		UpdatedDate: currentStr,
	}

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	mapperMock.On("MapDomainToResponse", domainUser).Return(responseUser)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	findByReferenceMock.On("Execute", reference).Return(domainUser, nil)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/USER1", nil)

	r := testRouter()
	r.GET("/api/v1/users/:id", handler.FindByReference)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result UserResponse
	json.NewDecoder(w.Body).Decode(&result)

	assert.NotNil(t, result)
	assert.Equal(t, responseUser, result)

	mapperMock.AssertExpectations(t)
	findByReferenceMock.AssertExpectations(t)
}

func TestUser_GivenAnId_WhenFindById_AndServiceReturnedAnError_ThenReturnInternalServerErrorResponse(t *testing.T) {
	t.Log("Failure to find an user by its id because service returned an error")

	reference := "USER1"

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	findByReferenceMock.On("Execute", reference).Return(domain.User{}, errors.New("service error"))
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/USER1", nil)

	r := testRouter()
	r.GET("/api/v1/users/:id", handler.FindByReference)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "service error", err.Message)

	findByReferenceMock.AssertExpectations(t)
}

func TestUser_GivenACreateRequest_WhenCreate_ThenReturnCreatedUserResponse(t *testing.T) {
	t.Log("Successfully create an user")

	current := time.Now().UTC()
	currentStr := current.Format(time.RFC3339)
	request := UserCreateRequest{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	domainInput := domain.UserCreateInput{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
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

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	mapperMock.On("MapCreateRequestToInput", request).Return(domainInput)
	mapperMock.On("MapDomainToResponse", domainUser).Return(responseUser)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	createMock.On("Execute", domainInput).Return(domainUser, nil)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(request)
	requestData := buffer.Bytes()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestData))

	r := testRouter()
	r.POST("/api/v1/users", handler.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var result UserResponse
	json.NewDecoder(w.Body).Decode(&result)

	assert.NotNil(t, result)
	assert.Equal(t, responseUser, result)

	mapperMock.AssertExpectations(t)
	createMock.AssertExpectations(t)
}

func TestUser_GivenANotValidCreateRequest_WhenCreate_ThenReturnBadRequestResponse(t *testing.T) {
	t.Log("Failure create an user because request has not a valid format")

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode("NOT A JSON")
	requestData := buffer.Bytes()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestData))

	r := testRouter()
	r.POST("/api/v1/users", handler.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "request body is not valid", err.Message)

	mapperMock.AssertExpectations(t)
	createMock.AssertExpectations(t)
}

func TestUser_GivenACreateRequestWithNotValidData_WhenCreate_ThenReturnBadRequestResponse(t *testing.T) {
	t.Log("Failure create an user because request has not a valid data")

	request := UserCreateRequest{
		FirstName: "Foo",
	}

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(request)
	requestData := buffer.Bytes()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestData))

	r := testRouter()
	r.POST("/api/v1/users", handler.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "request body is not valid", err.Message)

	mapperMock.AssertExpectations(t)
	createMock.AssertExpectations(t)
}

func TestUser_GivenACreateRequest_WhenCreate_AndServiceReturnedAnError_ThenReturnInternalServerErrorResponse(t *testing.T) {
	t.Log("Failure to create an user because service returned an error")

	request := UserCreateRequest{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	domainInput := domain.UserCreateInput{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	mapperMock.On("MapCreateRequestToInput", request).Return(domainInput)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	createMock.On("Execute", domainInput).Return(domain.User{}, errors.New("service error"))
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(request)
	requestData := buffer.Bytes()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestData))

	r := testRouter()
	r.POST("/api/v1/users", handler.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "service error", err.Message)

	mapperMock.AssertExpectations(t)
	createMock.AssertExpectations(t)
}

func TestUser_GivenAnUpdateRequest_WhenUpdate_ThenReturnUpdatedUserResponse(t *testing.T) {
	t.Log("Successfully update an user")

	current := time.Now().UTC()
	currentStr := current.Format(time.RFC3339)
	reference := "USER1"
	request := UserUpdateRequest{
		UserCreateRequest: UserCreateRequest{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
	}
	domainInput := domain.UserUpdateInput{
		Reference: reference,
		UserCreateInput: domain.UserCreateInput{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
	}
	domainUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference:   reference,
			IsActive:    true,
			CreatedDate: current,
			UpdatedDate: current,
		},
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	responseUser := UserResponse{
		Id:          reference,
		FirstName:   "Foo",
		LastName:    "Bar",
		Email:       "foobar@email.com",
		IsActive:    true,
		CreatedDate: currentStr,
		UpdatedDate: currentStr,
	}

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	mapperMock.On("MapUpdateRequestToInput", reference, request).Return(domainInput)
	mapperMock.On("MapDomainToResponse", domainUser).Return(responseUser)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	updateMock.On("Execute", domainInput).Return(domainUser, nil)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(request)
	requestData := buffer.Bytes()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/api/v1/users/USER1", bytes.NewBuffer(requestData))

	r := testRouter()
	r.PUT("/api/v1/users/:id", handler.Update)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result UserResponse
	json.NewDecoder(w.Body).Decode(&result)

	assert.NotNil(t, result)
	assert.Equal(t, responseUser, result)

	mapperMock.AssertExpectations(t)
	updateMock.AssertExpectations(t)
}

func TestUser_GivenANotValidUpdateRequest_WhenUpdate_ThenReturnBadRequestResponse(t *testing.T) {
	t.Log("Failure when update an user because request has not a valid format")

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode("NOT A JSON")
	requestData := buffer.Bytes()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/api/v1/users/USER1", bytes.NewBuffer(requestData))

	r := testRouter()
	r.PUT("/api/v1/users/:id", handler.Update)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "request body is not valid", err.Message)

	mapperMock.AssertExpectations(t)
	updateMock.AssertExpectations(t)
}

func TestUser_GivenAnUpdateRequestWithNotValidData_WhenUpdate_ThenReturnBadRequestResponse(t *testing.T) {
	t.Log("Failure when update an user because request has not a valid data")

	request := UserUpdateRequest{
		UserCreateRequest: UserCreateRequest{
			FirstName: "Foo",
		},
	}

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(request)
	requestData := buffer.Bytes()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/api/v1/users/USER1", bytes.NewBuffer(requestData))

	r := testRouter()
	r.PUT("/api/v1/users/:id", handler.Update)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "request body is not valid", err.Message)

	mapperMock.AssertExpectations(t)
	updateMock.AssertExpectations(t)
}

func TestUser_GivenAnUpdateRequest_WhenUpdate_AndServiceReturnedAnError_ThenReturnInternalServerErrorResponse(t *testing.T) {
	t.Log("Failure to update an user because service returned an error")

	reference := "USER1"
	request := UserUpdateRequest{
		UserCreateRequest: UserCreateRequest{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
	}
	domainInput := domain.UserUpdateInput{
		Reference: reference,
		UserCreateInput: domain.UserCreateInput{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foobar@email.com",
		},
	}

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	mapperMock.On("MapUpdateRequestToInput", reference, request).Return(domainInput)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	updateMock.On("Execute", domainInput).Return(domain.User{}, errors.New("service error"))
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(request)
	requestData := buffer.Bytes()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/api/v1/users/USER1", bytes.NewBuffer(requestData))

	r := testRouter()
	r.PUT("/api/v1/users/:id", handler.Update)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "service error", err.Message)

	mapperMock.AssertExpectations(t)
	updateMock.AssertExpectations(t)
}

func TestUser_GivenADeleteRequest_WhenDelete_ThenReturnDeletedUserResponse(t *testing.T) {
	t.Log("Successfully delete an user")

	current := time.Now().UTC()
	currentStr := current.Format(time.RFC3339)
	reference := "USER1"
	domainUser := domain.User{
		GenericEntity: domain.GenericEntity{
			Reference:   reference,
			IsActive:    false,
			CreatedDate: current,
			UpdatedDate: current,
		},
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	responseUser := UserResponse{
		Id:          reference,
		FirstName:   "Foo",
		LastName:    "Bar",
		Email:       "foobar@email.com",
		IsActive:    false,
		CreatedDate: currentStr,
		UpdatedDate: currentStr,
	}

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	mapperMock.On("MapDomainToResponse", domainUser).Return(responseUser)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	deleteMock.On("Execute", reference).Return(domainUser, nil)
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/users/USER1", nil)

	r := testRouter()
	r.DELETE("/api/v1/users/:id", handler.Delete)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result UserResponse
	json.NewDecoder(w.Body).Decode(&result)

	assert.NotNil(t, result)
	assert.Equal(t, responseUser, result)

	mapperMock.AssertExpectations(t)
	deleteMock.AssertExpectations(t)
}

func TestUser_GivenADeleteRequest_WhenDelete_AndServiceReturnedAnError_ThenReturnInternalServerErrorResponse(t *testing.T) {
	t.Log("Failure to delete an user because service returned an error")

	reference := "USER1"

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	deleteMock.On("Execute", reference).Return(domain.User{}, errors.New("service error"))
	searchMock := new(userSearchServiceMock)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/users/USER1", nil)

	r := testRouter()
	r.DELETE("/api/v1/users/:id", handler.Delete)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "service error", err.Message)

	mapperMock.AssertExpectations(t)
	deleteMock.AssertExpectations(t)
}

func TestUser_WhenSearch_ThenReturnSearchUserResponse(t *testing.T) {
	t.Log("Successfully search users")

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

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	mapperMock.On("MapDomainSearchOutputToResponse", domainSearchOutput).Return(searchResponse)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)
	searchMock.On("Execute", mock.AnythingOfType("UserSearchInput")).Return(domainSearchOutput, nil)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/search", nil)
	q := req.URL.Query()
	q.Add("page", fmt.Sprint(page))
	q.Add("size", fmt.Sprint(size))
	req.URL.RawQuery = q.Encode()

	r := testRouter()
	r.GET("/api/v1/users/search", handler.Search)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result UserSearchResponse
	json.NewDecoder(w.Body).Decode(&result)

	assert.NotNil(t, result)
	assert.Equal(t, searchResponse, result)

	mapperMock.AssertExpectations(t)
	searchMock.AssertExpectations(t)
}

func TestUser_WhenSearch_AndPageAndSizeAreNotSet_ThenReturnSearchUserResponseUsingDefaultValues(t *testing.T) {
	t.Log("Successfully search users")

	config := newApplicationConfigurationMock()
	total := int64(1)
	current := time.Now().UTC()
	currentStr := current.Format(time.RFC3339)
	domainSearchOutput := domain.UserSearchOutput{
		SearchOutput: domain.SearchOutput{
			Total:    total,
			Page:     config.PagingDefaultPage,
			PageSize: config.PagingDefaultSize,
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
		Page:     config.PagingDefaultPage,
		PageSize: config.PagingDefaultSize,
	}

	mapperMock := new(userMapperMock)
	mapperMock.On("MapDomainSearchOutputToResponse", domainSearchOutput).Return(searchResponse)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)
	searchMock.On("Execute", mock.AnythingOfType("UserSearchInput")).Return(domainSearchOutput, nil)

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/search", nil)

	r := testRouter()
	r.GET("/api/v1/users/search", handler.Search)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result UserSearchResponse
	json.NewDecoder(w.Body).Decode(&result)

	assert.NotNil(t, result)
	assert.Equal(t, searchResponse, result)

	mapperMock.AssertExpectations(t)
	searchMock.AssertExpectations(t)
}

func TestUser_WhenSearch_AndServiceFails_ThenReturnInternalServiceError(t *testing.T) {
	t.Log("Failure when search users because service returned an error")

	page := 1
	size := 10

	config := newApplicationConfigurationMock()
	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	searchMock := new(userSearchServiceMock)
	searchMock.On("Execute", mock.AnythingOfType("UserSearchInput")).Return(domain.UserSearchOutput{}, errors.New("service error"))

	handler := NewDefaultUser(config,
		mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock,
		searchMock)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/search", nil)
	q := req.URL.Query()
	q.Add("page", fmt.Sprint(page))
	q.Add("size", fmt.Sprint(size))
	req.URL.RawQuery = q.Encode()

	r := testRouter()
	r.GET("/api/v1/users/search", handler.Search)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var err libErrors.APIError
	json.NewDecoder(w.Body).Decode(&err)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "service error", err.Message)

	mapperMock.AssertExpectations(t)
	searchMock.AssertExpectations(t)
}
