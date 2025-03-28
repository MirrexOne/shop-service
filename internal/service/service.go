package service

import (
	"context"
	"shop-service/internal/repo"
	"shop-service/pkg/hasher"
	"time"
)

type AuthCreateUserInput struct {
	Username string
	Password string
}

type AuthGenerateTokenInput struct {
	Username string
	Password string
}

type SignInInput struct {
	Username string
	Password string
}

type Auth interface {
	CreateUser(ctx context.Context, input AuthCreateUserInput) (int, error)
	GenerateToken(ctx context.Context, input AuthGenerateTokenInput) (string, error)
	ParseToken(token string) (int, error)
}

type Services struct {
	Auth Auth
}

type ServicesDependencies struct {
	Repos  *repo.Repositories
	Hasher hasher.PasswordHasher

	SignKey  string
	TokenTTL time.Duration
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Auth: NewAuthService(deps.Repos.User, deps.Hasher, deps.SignKey, deps.TokenTTL),
	}
}
