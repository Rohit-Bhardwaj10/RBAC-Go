package repository

import (
	"context"

	model "github.com/Rohit-Bhardwaj10/RBAC-Go/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// save user detail
func (r *UserRepository) save(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// get all users
func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
