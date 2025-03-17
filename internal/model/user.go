package model

import "time"

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Coins     uint      `db:"coins"`
	CreatedAt time.Time `db:"created_at"`
}
