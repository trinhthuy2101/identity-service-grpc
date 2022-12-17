package grpc

import (
	"context"

	"github.com/uchin-mentorship/ecommerce-go/identity"

	"ecommerce/identity/internal/usecase"
)

type identityService struct {
	identity.IdentityServiceServer
	u usecase.AuthUsecase
}

var _ identity.IdentityServiceServer = (*identityService)(nil)

func (s *identityService) VerifyToken(ctx context.Context, req *identity.VerifyTokenRequest) (result *identity.VerifyTokenResponse, err error) {
	resp, err := s.u.VerifyToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &identity.VerifyTokenResponse{
		IsValid: true,
		User: &identity.AdminUser{
			Email: resp.Email,
		},
	}, err
}

func NewIdentityService(u usecase.AuthUsecase) identity.IdentityServiceServer {
	return &identityService{
		u: u,
	}
}
