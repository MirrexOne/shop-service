package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"shop-service/internal/model"
	"shop-service/internal/repo"
	"shop-service/internal/repo/repoerrs"
	"shop-service/pkg/hasher"
	"strings"
	"time"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type AuthService struct {
	userRepo       repo.User
	passwordHasher hasher.PasswordHasher
	signKey        string
	tokenTTL       time.Duration
}

func NewAuthService(userRepo repo.User, passwordHasher hasher.PasswordHasher, signKey string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		signKey:        signKey,
		tokenTTL:       tokenTTL,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, input AuthCreateUserInput) (int, error) {
	user := model.User{
		Username: input.Username,
		Password: s.passwordHasher.Hash(input.Password),
	}

	userId, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return 0, ErrUserAlreadyExists
		}
		log.Errorf("AuthService.CreateUser - c.userRepo.CreateUser: %v", err)
		return 0, ErrCannotCreateUser
	}

	return userId, nil
}

func (s *AuthService) GenerateToken(ctx context.Context, input AuthGenerateTokenInput) (string, error) {
	user, err := s.userRepo.GetUserByUsernameAndPassword(ctx, input.Username, s.passwordHasher.Hash(input.Password))
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrUserNotFound
		}
		log.Errorf("AuthService.GenerateToken: cannot get user: %v", err)
		return "", ErrCannotGetUser
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	tokenString, err := token.SignedString([]byte(s.signKey))
	if err != nil {
		log.Errorf("AuthService.GenerationToken: cannot sign token: %v", err)
		return "", ErrCannotSingToken
	}

	return tokenString, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.signKey), nil
	})

	if err != nil {
		return 0, ErrCannotParseToken
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, ErrCannotParseToken
	}

	return claims.UserId, nil
}

func (s *AuthService) SignInUser(ctx context.Context, username, password string) error {
	user, err := s.userRepo.GetUserByUsernameAndPassword(ctx, username, s.passwordHasher.Hash(password))
	if err != nil || errors.Is(err, repoerrs.ErrNotFound) {
		return ErrUserNotFound
	}

	isValidUsername := strings.EqualFold(username, user.Username)
	isEqual := s.compareHashAndPassword(user.Password, password)
	if !(isEqual || isValidUsername) {
		return ErrInvalidUsernameOrPassword
	}

	return nil
}

func (s *AuthService) compareHashAndPassword(hashedPassword, password string) bool {
	hash := s.passwordHasher.Hash(password)
	return strings.EqualFold(hash, hashedPassword)
}
