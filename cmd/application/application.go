package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"log"
	"thesis_back/internal/config"
)

type Application struct {
	config *config.Config
	db     *gorm.DB
	minio  *minio.Client
}

func NewApplication(config *config.Config, db *gorm.DB, minioClient *minio.Client) *Application {
	return &Application{
		config: config,
		db:     db,
		minio:  minioClient,
	}
}

func (a *Application) Start() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.Use(gin.Recovery())

	auth := v1.Group("/auth")
	{
		auth.POST("/register")
		auth.POST("/login")
		auth.POST("/refresh")
		auth.GET("/me")
	}

	projects := v1.Group("/projects")
	{
		projects.POST("")
		projects.GET("")
		projects.GET("/:id")
		projects.DELETE("/:id")
		projects.PUT("/:id")
	}

	layers := v1.Group("/layers")
	{
		layers.POST("")
		layers.GET("")
		layers.GET("/:id")
		layers.DELETE("/:id")
		layers.PUT("/:id")
	}

	images := v1.Group("/images")
	{
		images.GET("")
		images.POST("")
		images.PUT("")
		images.DELETE("")
	}

	if err := router.Run(fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port)); err != nil {
		log.Fatal(err)
	}
}
