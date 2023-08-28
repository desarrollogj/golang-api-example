package handler

import "github.com/gin-gonic/gin"

func testRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	return router
}
