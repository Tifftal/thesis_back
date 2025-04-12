package main

import (
	"fmt"
	"log"
	"thesis_back/cmd/application"
	"thesis_back/internal/config"
	"thesis_back/internal/infrastructure/db/postgres"
	"thesis_back/internal/infrastructure/s3/minio"
	"thesis_back/internal/pkg/logger"
	"thesis_back/internal/service"

	user_repo "thesis_back/internal/repository/user"
	user_handler "thesis_back/internal/transport/http/user"
	user_usecase "thesis_back/internal/usecase/user"

	project_repo "thesis_back/internal/repository/project"
	project_handler "thesis_back/internal/transport/http/project"
	project_usecase "thesis_back/internal/usecase/project"

	image_repo "thesis_back/internal/repository/image"
	image_handler "thesis_back/internal/transport/http/image"
	image_usecase "thesis_back/internal/usecase/image"

	layer_repo "thesis_back/internal/repository/layer"
	layer_handler "thesis_back/internal/transport/http/layer"
	layer_usecase "thesis_back/internal/usecase/layer"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("Server starting on %s:%d\n", cfg.Server.Host, cfg.Server.Port)

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:            cfg.DB.Host,
		Port:            cfg.DB.Port,
		User:            cfg.DB.User,
		Password:        cfg.DB.Password,
		DBName:          cfg.DB.DBName,
		SSLMode:         cfg.DB.SSLMode,
		MaxIdleConns:    cfg.DB.MaxIdleConns,
		MaxOpenConns:    cfg.DB.MaxOpenConns,
		ConnMaxLifetime: cfg.DB.ConnMaxLifetime,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := postgres.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	s3, err := minio.NewMinioClient(minio.Config{
		Endpoint:   cfg.S3.Endpoint,
		AccessKey:  cfg.S3.AccessKey,
		SecretKey:  cfg.S3.SecretKey,
		BucketName: cfg.S3.BucketName,
		UseSSL:     cfg.S3.UseSSL,
		Region:     cfg.S3.Region,
	})
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	custom_logger, err := logger.New(cfg.Logging)

	auth_service := service.NewAuthService(&service.JWTConfig{
		SecretKey:     cfg.Auth.JWTSecret,
		AccessExpiry:  cfg.Auth.AccessTokenExpire,
		RefreshExpiry: cfg.Auth.RefreshTokenExpire,
	})

	ur := user_repo.NewUserRepository(db)
	uc := user_usecase.NewUserUseCase(ur, auth_service, custom_logger)
	uh := user_handler.NewUserHandler(uc, custom_logger)

	pr := project_repo.NewProjectRepository(db)
	pc := project_usecase.NewProjectUseCase(&pr, custom_logger)
	ph := project_handler.NewProjectHandler(&pc, custom_logger)

	ir := image_repo.NewImageRepository(db, s3)
	ic := image_usecase.NewImageUseCase(&ir, custom_logger)
	ih := image_handler.NewImageHandler(&ic, custom_logger)

	lr := layer_repo.NewLayerRepository(db)
	lu := layer_usecase.NewLayerUseCase(&lr, custom_logger)
	lh := layer_handler.NewLayerHandler(&lu, custom_logger)

	app := application.NewApplication(cfg, custom_logger, db, s3)

	app.Start(uh, ph, lh, ih, auth_service)
}
