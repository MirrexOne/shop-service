package service

import "fmt"

var (
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrCannotCreateUser  = fmt.Errorf("cannot create user")
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrCannotGetUser     = fmt.Errorf("cannot get user")

	ErrCannotSingToken  = fmt.Errorf("cannot sign token")
	ErrCannotParseToken = fmt.Errorf("cannot parse token")

	ErrInvalidPassword = fmt.Errorf("invalid password")
)
