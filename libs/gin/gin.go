package gin

import (
	"fmt"
	"strconv"

	"github.com/desarrollogj/golang-api-example/libs/errors"
	"github.com/gin-gonic/gin"
)

// WrapperFunc is the func type for the custom handlers.
type WrapperFunc func(c *gin.Context) *errors.APIError

// ErrorWrapper if handlerFunc returns an error, then the response object will be composed from the error's information.
func ErrorWrapper(handlerFunc WrapperFunc, c *gin.Context) {
	err := handlerFunc(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

// NoRouteHandler handles requests for non registered routes
func NoRouteHandler(c *gin.Context) {
	ErrorWrapper(func(c *gin.Context) *errors.APIError {
		return errors.NewResourceNotFound(fmt.Sprintf("Resource not found for %s.", c.Request.URL.Path))
	}, c)
}

// MethodNotAllowedHandler handles requests for registered routes with invalid http methods on their requests
func MethodNotAllowedHandler(c *gin.Context) {
	ErrorWrapper(func(c *gin.Context) *errors.APIError {
		return errors.NewMethodNotAllowed("Method not allowed - %s - %s ", c.Request.Method, c.Request.URL.Path)
	}, c)
}

// GetIntQuery recovers an integer value from the querystring
func GetIntQuery(key string, c *gin.Context) int {
	value := c.Query(key)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		intValue = 0
	}
	return intValue
}
