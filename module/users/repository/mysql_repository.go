package repository

import (
	"database/sql"
	"fmt"
	"hacktiv8/final/helper"
	"hacktiv8/final/module/users"
	"hacktiv8/final/module/users/model"
	"strconv"
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

	// Photo Profile step

	defaultPhotoProfile := "https://i.ibb.co/whHn3mf/default-photo-profile.png"

	if u.PhotoProfile != "" {
		imgBB := helper.NewImgBBConn()
		imgURL, err := imgBB.Upload(u.PhotoProfile)

		if err != nil {
			return err
		}

		u.PhotoProfile = imgURL
	} else {
		u.PhotoProfile = defaultPhotoProfile
	}

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

	result, err := statement.Exec(u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.PhotoProfile, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		return err
	}

	lastInsertIdInt64, err := result.LastInsertId()

	if err != nil {
		return err
	}

	lastInsertIdStr := strconv.FormatInt(lastInsertIdInt64, 10)

	u.UserID = lastInsertIdStr

	return nil
}

// Update User
func (r *mysqlUsersRepository) Update(id string, u *model.User) (rowAffected *string, err error) {

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
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.PhotoProfile, u.UpdatedAt, id)

	if err != nil {
		return nil, err
	}

	rowsAffectedInt64, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	rowsAffectedStr := strconv.FormatInt(rowsAffectedInt64, 10)

	return &rowsAffectedStr, nil
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
