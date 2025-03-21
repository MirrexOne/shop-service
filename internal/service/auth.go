package service

import (
	"context"
	"shop-service/internal/model"
	"shop-service/internal/repo"
	"shop-service/pkg/hasher"
)

type AuthService struct {
	userRepo       repo.User
	passwordHasher hasher.PasswordHasher
}

func NewAuthService(userRepo repo.User, passwordHasher hasher.PasswordHasher) *AuthService {
	return &AuthService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, input AuthCreateUserInput) error {
	user := model.User{
		Username: input.Username,
		Password: s.passwordHasher.Hash(input.Password),
	}

	err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
