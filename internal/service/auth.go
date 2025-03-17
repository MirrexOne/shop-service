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

func (s *AuthService) CreateUser(ctx context.Context, input AuthCreateUserInput) (int, error) {
	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	userId, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *AuthService) GetUserById(ctx context.Context, id int) (*model.User, error) {
	return s.userRepo.GetUserById(ctx, id)
}
