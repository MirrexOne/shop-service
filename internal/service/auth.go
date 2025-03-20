package service

import (
	"context"
	"shop-service/internal/model"
	"shop-service/internal/repo"
)

type AuthService struct {
	userRepo repo.User
}

func NewAuthService(userRepo repo.User) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) CreateUser(ctx context.Context, input AuthCreateUserInput) error {
	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
