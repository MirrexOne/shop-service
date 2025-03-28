package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shop-service/internal/model"
	"shop-service/internal/repo/repoerrs"
)

type UserRepo struct {
	*sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser(ctx context.Context, user model.User) (int, error) {
	const op = "repo.db.CreateUser"

	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return 0, fmt.Errorf("%s: begin transaction %w", op, err)
	}

	defer func() {
		var e error
		if err == nil {
			e = tx.Commit()
		} else {
			e = tx.Rollback()
		}

		if err == nil && e != nil {
			err = fmt.Errorf("finishing transaction: %w", e)
		}
	}()

	var userId int
	createUser := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING ID"
	err = tx.QueryRowContext(ctx, createUser, user.Username, user.Password).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("creating user: %w", err)
	}

	createWallet := "INSERT INTO wallet (user_id) VALUES ($1)"
	_, err = tx.ExecContext(ctx, createWallet, userId)
	if err != nil {
		return 0, fmt.Errorf("creating wallet: %w", err)
	}

	return userId, nil
}

func (r *UserRepo) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (model.User, error) {
	query := `
	SELECT u
	FROM users u
	WHERE username = $1 AND password = $2
	`

	var user model.User

	err := r.QueryRowContext(ctx, query, username, password).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repoerrs.ErrNotFound
		}
		return model.User{}, fmt.Errorf("UserRepo.GetUserByUsernameAndPassword - r.QueryRowContext: %v", err)
	}

	return user, nil
}
