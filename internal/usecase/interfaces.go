// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"ecommerce/identity/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	AuthUsecase interface {
		Login(ctx context.Context, request *entity.AdminUser) (*entity.AdminUser, string, error)
		Register(ctx context.Context, request *entity.AdminUser) error
		VerifyToken(ctx context.Context, tokenString string) (*entity.AdminUser, error)
		// Get(ctx context.Context,id uint32)(*entity.AdminUser,error)
	}

	AuthRepo interface {
		Login(ctx context.Context, request *entity.AdminUser) (*entity.AdminUser, error)
		Register(ctx context.Context, request *entity.AdminUser) error
		Get(ctx context.Context, id uint32) (*entity.AdminUser, error)
	}
)
