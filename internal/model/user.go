package model

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID       int64     `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	CreateAt time.Time `db:"created_at"`
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return errors.New("Email must exists '@' and '.'")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 || !strings.ContainsAny("0123456789", password) || !strings.Contains("_", password) {
		return errors.New("Password must exist '_', number and length > 7")
	}
	return nil
}
