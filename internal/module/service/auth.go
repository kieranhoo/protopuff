package service

import (
	"context"
	"protopuff/internal/proto/gen/v1/auth"
)

type Auth struct {
	auth.UnimplementedAuthServer
}

func NewAuth() auth.AuthServer {
	return &Auth{}
}

func (a *Auth) SignIn(ctx context.Context, data *auth.SignInRequest) (*auth.SignInResponse, error) {
	return &auth.SignInResponse{
		Success:     true,
		Email:       data.Email,
		AccessToken: "0x0",
	}, nil
}
