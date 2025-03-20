package repo

import (
	"context"
	"database/sql"
	"shop-service/internal/model"
	"shop-service/internal/repo/db"
)

type User interface {
	CreateUser(ctx context.Context, user model.User) error
}

type Repositories struct {
	User
}

func NewRepositories(pgdb *sql.DB) *Repositories {
	return &Repositories{
		User: db.NewUserRepo(pgdb),
	}
}
