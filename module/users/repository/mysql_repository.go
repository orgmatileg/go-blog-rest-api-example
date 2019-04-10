package repository

import (
	"database/sql"
	"fmt"
	"hacktiv8/final/module/users"
	"hacktiv8/final/module/users/model"
	"time"
)

// mysqlUsersRepository struct
type mysqlUsersRepository struct {
	db *sql.DB
}

// NewUserRepositoryMysql NewUserRepositoryMysql
func NewUserRepositoryMysql(db *sql.DB) users.Repository {
	return &mysqlUsersRepository{db}
}

// Save User
func (r *mysqlUsersRepository) Save(u *model.User) error {

	query := `INSERT INTO tbl_users (
	username,
	email,
	password,
	first_name,
	last_name,
	photo_profile,
	created_at,
	updated_at)
	VALUES (?,?,?,?,?,?,?,?)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.PhotoProfile, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// Update User
func (r *mysqlUsersRepository) Update(id string, u *model.User) error {

	query := `
	UPDATE tbl_users SET 
	username=?, 
	email=?, 
	password=?,
	first_name=?,
	last_name=?,
	photo_profile=?,
	updated_at=?
	WHERE user_id=?
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.PhotoProfile, u.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil
}

// Delete User
func (r *mysqlUsersRepository) Delete(id string) error {

	query := `
	UPDATE tbl_users SET 
	deleted_at = ?
	WHERE user_id = ?`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(time.Now(), id)

	if err != nil {
		return err
	}

	return nil
}

// FindByID User
func (r *mysqlUsersRepository) FindByID(id string) (*model.User, error) {

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
	FROM tbl_users WHERE user_id = ? && deleted_at is NULL`

	var user model.User

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll User
func (r *mysqlUsersRepository) FindAll(limit, offset, order string) (model.Users, error) {

	query := fmt.Sprintf(`
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
	WHERE deleted_at is NULL
	ORDER BY created_at %s
	LIMIT %s 
	OFFSET %s`, order, limit, offset)

	var users model.Users

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User

		err = rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
