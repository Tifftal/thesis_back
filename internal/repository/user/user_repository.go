package user

import (
	"context"
	"gorm.io/gorm"
	"thesis_back/internal/domain"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	err := r.db.Delete(&domain.User{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
