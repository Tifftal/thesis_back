package application

import (
	"fmt"
	"log"
	_ "thesis_back/docs"
	"thesis_back/internal/config"
	"thesis_back/internal/service"
	image_handler "thesis_back/internal/transport/http/image"
	layer_handler "thesis_back/internal/transport/http/layer"
	"thesis_back/internal/transport/http/middleware"
	project_handler "thesis_back/internal/transport/http/project"
	user_handler "thesis_back/internal/transport/http/user"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func (a *Application) Start(user_handler *user_handler.UserHandler, project_handler *project_handler.ProjectHandler, layer_handler *layer_handler.LayerHandler, image_handler *image_handler.ImageHandler, auth_service *service.AuthService) {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Recovery())

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", user_handler.Register)
			auth.POST("/login", user_handler.Login)
			auth.POST("/refresh", user_handler.Refresh)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.IsAuthenticated(auth_service, a.logger.Named("Auth Middleware")))
		{
			user := protected.Group("/user")
			{
				user.GET("/me", user_handler.Me)
			}

			project := protected.Group("/project")
			{
				project.POST("", project_handler.Create)
				project.GET("", project_handler.Get)
				project.GET("/:id", project_handler.GetByID)
				project.DELETE("/:id", project_handler.Delete)
				project.PUT("/:id", project_handler.Update)
			}

			image := protected.Group("/image")
			{
				image.POST("", image_handler.Create)
				image.DELETE("/:id", image_handler.Delete)
				image.PUT("/:id", image_handler.Update)
			}

			layer := protected.Group("/layer")
			{
				layer.POST("", layer_handler.Create)
				layer.DELETE("/:id", layer_handler.Delete)
				layer.PUT("/:id", layer_handler.Update)
			}
		}
	}

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port)
	a.logger.Info("Starting server", zap.String("address", serverAddr))
	if err := router.Run(serverAddr); err != nil {
		a.logger.Fatal("Failed to start server", zap.Error(err))
		log.Fatal(err)
	}
}
