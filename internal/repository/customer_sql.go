package repository

import (
	"context"

	"gorm.io/gorm"

	"ecommerce/identity/internal/entity"
)

type AuthRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (r *AuthRepo) withContext(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&entity.AdminUser{})
}

func (u *AuthRepo) Login(request *entity.AdminUser) (*entity.AdminUser, error) {
	return nil, nil
}

func (u *AuthRepo) Register(request *entity.AdminUser) error {
	return nil
}
func (u *AuthRepo) Get(ctx context.Context, id uint32) (*entity.AdminUser, error) {
	return nil, nil
}
