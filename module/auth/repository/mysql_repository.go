package repository

import (
	"database/sql"
	"hacktiv8/final/module/auth"
	"hacktiv8/final/module/users/model"
)

// mysqlAuthRepository struct
type mysqlAuthRepository struct {
	db *sql.DB
}

// NewUserRepositoryMysql NewUserRepositoryMysql
func NewAuthRepositoryMysql(db *sql.DB) auth.Repository {
	return &mysqlAuthRepository{db}
}

// Login
func (r *mysqlAuthRepository) Login(u *model.User) (user *model.User, err error) {

	query := `
		SELECT  
		user_id,
		username,
		email,
		password,
		first_name,
		last_name,
		photo_profile,
		created_at,
		updated_at 
		FROM tbl_users 
		WHERE username = ? or email = ?
		LIMIT 1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	var mu model.User

	err = statement.QueryRow(u.Username, u.Email).Scan(&mu.UserID, &mu.Username, &mu.Email, &mu.Password, &mu.FirstName, &mu.LastName, &mu.PhotoProfile, &mu.CreatedAt, &mu.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &mu, err
}
