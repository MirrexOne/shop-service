package db

import (
	"context"
	"database/sql"
	"fmt"
	"shop-service/internal/model"
)

type UserRepo struct {
	*sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser(ctx context.Context, user model.User) (int, error) {
	const op = "repo.db.CreateUser"
	var id int

	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return 0, fmt.Errorf("%s: begin transaction %w", op, err)
	}

	createUser := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING ID"
	_, err = tx.ExecContext(ctx, createUser, user.Username, user.Password)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: execute create user %w", op, err)
	}

	createWallet := "INSERT INTO wallet (user_id) VALUES ($1)"
	_, err = tx.ExecContext(ctx, createWallet, id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: execute create wallet %w", op, err)
	}

	return id, nil
}

func (r *UserRepo) GetUserById(ctx context.Context, id int) (*model.User, error) {
	createStmt := "SELECT * FROM users WHERE id = ?"

	var user model.User

	err := r.QueryRowContext(ctx, createStmt, id).Scan(&user)
	if err != nil {
		return &user, err
	}

	return &user, nil
}
