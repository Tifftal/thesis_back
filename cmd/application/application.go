package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	_ "thesis_back/docs"
	"thesis_back/internal/config"
	"thesis_back/internal/service"
	"thesis_back/internal/transport/http/middleware"
	project_handler "thesis_back/internal/transport/http/project"
	user_handler "thesis_back/internal/transport/http/user"
)

type Application struct {
	config *config.Config
	logger *zap.Logger
	db     *gorm.DB
	minio  *minio.Client
}

func NewApplication(config *config.Config, logger *zap.Logger, db *gorm.DB, minioClient *minio.Client) *Application {
	return &Application{
		config: config,
		logger: logger,
		db:     db,
		minio:  minioClient,
	}
}

// Start @title Thesis Backend API
// @version 1.0
// @description API для дипломного проекта
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
func (a *Application) Start(user_handler *user_handler.UserHandler, project_handler *project_handler.ProjectHandler, auth_service *service.AuthService) {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := v1.Group("/auth")
	{
		auth.POST("/register", user_handler.Register)
		auth.POST("/login", user_handler.Login)
		auth.POST("/refresh", user_handler.Refresh)
	}

	protected := v1.Group("")
	protected.Use(middleware.IsAuthenticated(auth_service, a.logger.Named("Auth Middleware")))
	{
		user := protected.Group("/user")
		{
			user.GET("/me", user_handler.Me)
		}

		projects := protected.Group("/projects")
		{
			projects.POST("", project_handler.Create)
			projects.GET("", project_handler.Get)
			projects.GET("/:id", project_handler.GetByID)
			projects.DELETE("/:id", project_handler.Delete)
			projects.PUT("/:id", project_handler.Update)
		}

		layers := protected.Group("/layers")
		{
			layers.POST("")
			layers.GET("")
			layers.GET("/:id")
			layers.DELETE("/:id")
			layers.PUT("/:id")
		}

		images := protected.Group("/images")
		{
			images.GET("")
			images.POST("")
			images.PUT("")
			images.DELETE("")
		}
	}

	if err := router.Run(fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port)); err != nil {
		log.Fatal(err)
	}
}
