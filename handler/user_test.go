package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/desarrollogj/golang-api-example/domain"
	libErrors "github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/stretchr/testify/assert"
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

	mapperMock := new(userMapperMock)
	mapperMock.On("MapDomainListToResponseList", domainUsers).Return(responseUsers)
	findAllMock := new(userFindAllServiceMock)
	findAllMock.On("Execute").Return(domainUsers, nil)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findAllMock.On("Execute").Return([]domain.User{}, errors.New("service error"))
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	mapperMock.On("MapDomainToResponse", domainUser).Return(responseUser)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	findByReferenceMock.On("Execute", reference).Return(domainUser, nil)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	findByReferenceMock.On("Execute", reference).Return(domain.User{}, errors.New("service error"))
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	mapperMock.On("MapCreateRequestToInput", request).Return(domainInput)
	mapperMock.On("MapDomainToResponse", domainUser).Return(responseUser)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	createMock.On("Execute", domainInput).Return(domainUser, nil)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	mapperMock.On("MapCreateRequestToInput", request).Return(domainInput)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	createMock.On("Execute", domainInput).Return(domain.User{}, errors.New("service error"))
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	mapperMock.On("MapUpdateRequestToInput", reference, request).Return(domainInput)
	mapperMock.On("MapDomainToResponse", domainUser).Return(responseUser)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	updateMock.On("Execute", domainInput).Return(domainUser, nil)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	mapperMock.On("MapUpdateRequestToInput", reference, request).Return(domainInput)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	updateMock.On("Execute", domainInput).Return(domain.User{}, errors.New("service error"))
	deleteMock := new(userDeleteServiceMock)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	mapperMock.On("MapDomainToResponse", domainUser).Return(responseUser)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	deleteMock.On("Execute", reference).Return(domainUser, nil)

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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

	mapperMock := new(userMapperMock)
	findAllMock := new(userFindAllServiceMock)
	findByReferenceMock := new(userFindByReferenceServiceMock)
	createMock := new(userCreateServiceMock)
	updateMock := new(userUpdateServiceMock)
	deleteMock := new(userDeleteServiceMock)
	deleteMock.On("Execute", reference).Return(domain.User{}, errors.New("service error"))

	handler := NewDefaultUser(mapperMock,
		findAllMock,
		findByReferenceMock,
		createMock,
		updateMock,
		deleteMock)

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
