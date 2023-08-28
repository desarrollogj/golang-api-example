package main

import (
	"fmt"

	_ "github.com/desarrollogj/golang-api-example/docs"
	"github.com/desarrollogj/golang-api-example/libs/database"
	"github.com/desarrollogj/golang-api-example/libs/logger"
	"github.com/desarrollogj/golang-api-example/libs/system"
	"github.com/desarrollogj/golang-api-example/router"
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
)

// @title           Users example api
// @version         0.0.1
// @description     A CRUD example api using Go language
// @BasePath  		/api/v1/users
func main() {
	// Defer database disconnection (if database client was created)
	defer database.MongoDisconnect()

	// Load configuration
	config.WithOptions(config.ParseEnv)
	config.AddDriver(json.Driver)

	env := system.GetEnv("APP_PROFILE", "LOCAL")
	confFile := fmt.Sprintf("config/config-%s.json", env)
	err := config.LoadFiles(confFile)
	if err != nil {
		fmt.Printf("Error while loading configuration file")
		panic(err)
	}

	// Load logger
	logger.InitLogger()
	logger.AppLog.Debug().Msg("Logger initialized")
	logger.AppLog.Info().Msg(fmt.Sprintf("Current environment: %s", env))

	// Load database
	database.Mongo = database.MongoConnect()

	// Create HTTP router and start
	gin.SetMode(config.String("ginMode", "debug"))
	r := router.CreateRouter()
	port := config.String("port")
	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("error while starting application")
		panic(err)
	}
}
