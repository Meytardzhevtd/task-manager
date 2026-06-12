package repository

import (
	"database/sql"
	"task-manager/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(username, email, password string) (*model.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	user := &model.User{Username: username, Email: email, Password: password}
	err := r.db.QueryRow(query, username, email, password).Scan(&user.ID, &user.CreateAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
