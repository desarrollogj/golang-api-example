package handler

import (
	"net/http"

	"github.com/desarrollogj/golang-api-example/domain"
	appErrors "github.com/desarrollogj/golang-api-example/libs/errors"
	appGin "github.com/desarrollogj/golang-api-example/libs/gin"
	"github.com/desarrollogj/golang-api-example/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// User represents the method for user endpoints handlers
type User interface {
	FindAll(c *gin.Context)
	FindByReference(c *gin.Context)
	Search(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// defaultUser is the default implementation for User interface
type defaultUser struct {
	config          domain.ApplicationConfiguration
	mapper          UserMapper
	findAll         user.FindAll
	findByReference user.FindByReference
	create          user.Create
	update          user.Update
	delete          user.Delete
	search          user.Search
}

// NewDefaultUser creates a defaultUser handler
func NewDefaultUser(config domain.ApplicationConfiguration,
	mapper UserMapper,
	findAll user.FindAll,
	findByReference user.FindByReference,
	create user.Create,
	update user.Update,
	delete user.Delete,
	search user.Search) defaultUser {
	validate = validator.New()
	return defaultUser{
		config:          config,
		mapper:          mapper,
		findAll:         findAll,
		findByReference: findByReference,
		create:          create,
		update:          update,
		delete:          delete,
		search:          search}
}

// FindAll find all users
// @Tags user
// @Summary Find all users
// @Description Find all users
// @Produce json
// @Success 200 {object} []handler.UserResponse
// @Failure 400	{object} appErrors.APIError
// @Failure 404	{object} appErrors.APIError
// @Failure 500	{object} appErrors.APIError
// @Router / [get]
func (h defaultUser) FindAll(c *gin.Context) {
	appGin.ErrorWrapper(h.executeFindAll, c)
}

func (h defaultUser) executeFindAll(c *gin.Context) *appErrors.APIError {
	users, err := h.findAll.Execute()
	if err != nil {
		return appErrors.HandleBusinessError(err)
	}

	c.JSON(http.StatusOK, h.mapper.MapDomainListToResponseList(users))
	return nil
}

// FindByReference find an user by its id
// @Tags user
// @Summary Find an user by its id
// @Description Find an user by its id
// @Param id path string true "User id"
// @Produce json
// @Success 200 {object} handler.UserResponse
// @Failure 400	{object} appErrors.APIError
// @Failure 404	{object} appErrors.APIError
// @Failure 500	{object} appErrors.APIError
// @Router /{id} [get]
func (h defaultUser) FindByReference(c *gin.Context) {
	appGin.ErrorWrapper(h.executeFindByReference, c)
}

func (h defaultUser) executeFindByReference(c *gin.Context) *appErrors.APIError {
	reference := c.Param("id")
	if len(reference) == 0 {
		return appErrors.NewBadRequest("user id is required")
	}

	user, err := h.findByReference.Execute(reference)
	if err != nil {
		return appErrors.HandleBusinessError(err)
	}

	c.JSON(http.StatusOK, h.mapper.MapDomainToResponse(user))
	return nil
}

// Search search users
// @Tags user
// @Summary Search users
// @Description Search users
// @Param firstName query string false "User first name"
// @Param lastName query string false "User last name"
// @Param email query string false "User email"
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Produce json
// @Success 200 {object} handler.UserSearchResponse
// @Failure 400	{object} appErrors.APIError
// @Failure 404	{object} appErrors.APIError
// @Failure 500	{object} appErrors.APIError
// @Router /search [get]
func (h defaultUser) Search(c *gin.Context) {
	appGin.ErrorWrapper(h.executeSearch, c)
}

func (h defaultUser) executeSearch(c *gin.Context) *appErrors.APIError {
	firstName := c.Query("firstName")
	lastName := c.Query("lastName")
	email := c.Query("email")
	page := appGin.GetIntQuery("page", c)
	size := appGin.GetIntQuery("size", c)
	if page < 1 {
		page = h.config.PagingDefaultPage
	}
	if size < 1 {
		size = h.config.PagingDefaultSize
	}

	input := domain.UserSearchInput{
		SearchInput: domain.SearchInput{
			Page:     page,
			PageSize: size,
		},
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	output, err := h.search.Execute(input)
	if err != nil {
		return appErrors.HandleBusinessError(err)
	}

	c.JSON(http.StatusOK, h.mapper.MapDomainSearchOutputToResponse(output))
	return nil
}

// Create creates an user
// @Tags user
// @Summary Create an user
// @Description Create an user
// @Param request body handler.UserCreateRequest true "user data"
// @Produce json
// @Success 201 {object} handler.UserResponse
// @Failure 400	{object} appErrors.APIError
// @Failure 404	{object} appErrors.APIError
// @Failure 500	{object} appErrors.APIError
// @Router / [post]
func (h defaultUser) Create(c *gin.Context) {
	appGin.ErrorWrapper(h.executeCreate, c)
}

func (h defaultUser) executeCreate(c *gin.Context) *appErrors.APIError {
	var req UserCreateRequest
	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		return appErrors.NewBadRequest("request body is not valid")
	}

	err = validate.Struct(req)
	if err != nil {
		return appErrors.NewBadRequest("request body is not valid")
	}

	created, err := h.create.Execute(h.mapper.MapCreateRequestToInput(req))
	if err != nil {
		return appErrors.HandleBusinessError(err)
	}

	c.JSON(http.StatusCreated, h.mapper.MapDomainToResponse(created))
	return nil
}

// Update update an user
// @Tags user
// @Summary Update an user
// @Description Update an user
// @Param id path string true "User id"
// @Param request body handler.UserUpdateRequest true "user data"
// @Produce json
// @Success 200 {object} handler.UserResponse
// @Failure 400	{object} appErrors.APIError
// @Failure 404	{object} appErrors.APIError
// @Failure 500	{object} appErrors.APIError
// @Router /{id} [put]
func (h defaultUser) Update(c *gin.Context) {
	appGin.ErrorWrapper(h.executeUpdate, c)
}

func (h defaultUser) executeUpdate(c *gin.Context) *appErrors.APIError {
	reference := c.Param("id")
	if len(reference) == 0 {
		return appErrors.NewBadRequest("user id is required")
	}
	var req UserUpdateRequest
	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		return appErrors.NewBadRequest("request body is not valid")
	}

	err = validate.Struct(req)
	if err != nil {
		return appErrors.NewBadRequest("request body is not valid")
	}

	updated, err := h.update.Execute(h.mapper.MapUpdateRequestToInput(reference, req))
	if err != nil {
		return appErrors.HandleBusinessError(err)
	}

	c.JSON(http.StatusOK, h.mapper.MapDomainToResponse(updated))
	return nil
}

// Delete delete an user
// @Tags user
// @Summary Delete an user
// @Description Delete an user
// @Param id path string true "User id"
// @Produce json
// @Success 200 {object} handler.UserResponse
// @Failure 400	{object} appErrors.APIError
// @Failure 404	{object} appErrors.APIError
// @Failure 500	{object} appErrors.APIError
// @Router /{id} [delete]
func (h defaultUser) Delete(c *gin.Context) {
	appGin.ErrorWrapper(h.executeDelete, c)
}

// Delete an user
func (h defaultUser) executeDelete(c *gin.Context) *appErrors.APIError {
	reference := c.Param("id")
	if len(reference) == 0 {
		return appErrors.NewBadRequest("user id is required")
	}

	deleted, err := h.delete.Execute(reference)
	if err != nil {
		return appErrors.HandleBusinessError(err)
	}

	c.JSON(http.StatusOK, h.mapper.MapDomainToResponse(deleted))
	return nil
}
