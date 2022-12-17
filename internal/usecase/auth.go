package usecase

import (
	"context"
	"ecommerce/identity/internal/entity"
	repo "ecommerce/identity/internal/repository"
	service "ecommerce/identity/internal/service/jwthelper"
)

type authUsecase struct {
	repo       repo.AuthRepo
	jwtService service.JWTHelper
}

func NewAuthUsecase(repo repo.AuthRepo, jwtService service.JWTHelper) AuthUsecase {
	return &authUsecase{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (u *authUsecase) Login(ctx context.Context, request *entity.AdminUser) (*entity.AdminUser, string, error) {
	user, err := u.repo.Login(request)
	if err != nil || user.ID == 0 {
		return nil, "", err
	}

	var token string

	token, err = u.jwtService.GenerateJWT(request.Email)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (u *authUsecase) Register(ctx context.Context, request *entity.AdminUser) error {
	return u.repo.Register(request)
}

func (u *authUsecase) VerifyToken(ctx context.Context, tokenString string) (*entity.AdminUser, error) {
	claims, err := u.jwtService.ValidateJWT(tokenString)
	if err != nil {
		return nil, err
	}

	return u.repo.Get(ctx, claims.ID)
}
