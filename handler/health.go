package handler

import (
	"net/http"

	"github.com/desarrollogj/golang-api-example/libs/system"
	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, struct {
		Status      string `json:"status"`
		Environment string `json:"environment"`
		App         string `json:"app"`
		Version     string `json:"version"`
	}{
		"OK",
		system.GetEnv("APP_PROFILE", "LOCAL"),
		system.GetEnv("APP_ARTIFACT", "UNKNOWN"),
		system.GetEnv("APP_VERSION", "UNKNOWN"),
	})
}
