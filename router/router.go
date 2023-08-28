package router

import (
	"github.com/desarrollogj/golang-api-example/domain"
	"github.com/desarrollogj/golang-api-example/handler"
	"github.com/desarrollogj/golang-api-example/infrastructure"
	appGin "github.com/desarrollogj/golang-api-example/libs/gin"
	"github.com/desarrollogj/golang-api-example/libs/logger"
	"github.com/desarrollogj/golang-api-example/user"
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	swagFiles "github.com/swaggo/files"
	swagGin "github.com/swaggo/gin-swagger"
)

// CreateRouter creates a GIN router
func CreateRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery(), logger.GinCustomLogger())

	router.HandleMethodNotAllowed = true

	router.NoMethod(appGin.MethodNotAllowedHandler)
	router.NoRoute(appGin.NoRouteHandler)

	// Swagger
	router.GET("/docs/*any", swagGin.WrapHandler(swagFiles.Handler))

	// Load api dependencies and mappings
	configureMappings(router)

	return router
}

// configureMappings configures Api endpoints
func configureMappings(router *gin.Engine) {
	// Configurations
	mongoRepoConfig := domain.MongoRepositoryConfiguration{}
	err := config.BindStruct("database", &mongoRepoConfig)
	if err != nil {
		logger.AppLog.Fatal().Err(err).Msg("unable to load repository configuration")
	}

	// Infrastructure
	userMongoRepositoryMapper := infrastructure.NewDefaultMongoRepositoryMapper()
	userMongoRepository := infrastructure.NewMongoUserRepository(mongoRepoConfig, userMongoRepositoryMapper)

	// Services
	userFindAllUC := user.NewDefaultFindAll(userMongoRepository)
	userFindByReferenceUC := user.NewDefaultFindByReference(userMongoRepository)
	userCreateUC := user.NewDefaultCreate(userMongoRepository)
	userUpdateUC := user.NewDefaultUpdate(userMongoRepository)
	userDeleteUC := user.NewDefaultDelete(userMongoRepository)

	// Handlers
	userMapper := handler.NewDefaultUserMapper()
	userHandler := handler.NewDefaultUser(userMapper,
		userFindAllUC,
		userFindByReferenceUC,
		userCreateUC,
		userUpdateUC,
		userDeleteUC)

	// Routes
	router.GET("/health", handler.Health)

	api := router.Group("/api/v1")
	api.GET("/users", userHandler.FindAll)
	api.GET("/users/:id", userHandler.FindByReference)
	api.POST("/users", userHandler.Create)
	api.PUT("/users/:id", userHandler.Update)
	api.DELETE("/users/:id", userHandler.Delete)
}
