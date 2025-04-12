package main

import (
	"fmt"
	"log"
	"thesis_back/cmd/application"
	"thesis_back/internal/config"
	"thesis_back/internal/infrastructure/db/postgres"
	"thesis_back/internal/infrastructure/s3/minio"
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

	minioClient, err := minio.NewMinioClient(minio.Config{
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

	app := application.NewApplication(cfg, db, minioClient)

	app.Start()
}
