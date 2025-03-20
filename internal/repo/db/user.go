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

func (r *UserRepo) CreateUser(ctx context.Context, user model.User) error {
	const op = "repo.db.CreateUser"

	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("%s: begin transaction %w", op, err)
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
		return fmt.Errorf("creating user: %w", err)
	}

	createWallet := "INSERT INTO wallet (user_id) VALUES ($1)"
	_, err = tx.ExecContext(ctx, createWallet, userId)
	if err != nil {
		return fmt.Errorf("creating wallet: %w", err)
	}

	return nil
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
