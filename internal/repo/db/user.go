package db

import (
	"context"
	"database/sql"
	"shop-service/internal/model"
)

type UserRepo struct {
	*sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser(ctx context.Context, user model.User) (int, error) {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	createStmt := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING ID"

	var id int
	stmt, err := tx.PrepareContext(ctx, createStmt)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = tx.QueryRowContext(ctx, createStmt, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	_, err = tx.ExecContext(ctx, ""+
		"INSERT INTO wallet (user_id) VALUES ($1)", id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetUserById(ctx context.Context, id int) (*model.User, error) {
	createStmt := "SELECT * FROM users WHERE id = ?"

	var user model.User

	stmt, err := r.DB.PrepareContext(ctx, createStmt)
	if err != nil {
		return &user, err
	}

	err = r.QueryRowContext(ctx, createStmt, id).Scan(&user)
	if err != nil {
		return &user, err
	}
	defer stmt.Close()

	return &user, nil
}
