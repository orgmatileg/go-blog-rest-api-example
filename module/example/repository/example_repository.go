package repository

import (
	"database/sql"
	"fmt"
	"hacktiv8/final/module/example"
	"hacktiv8/final/module/example/model"
	"strconv"
	"time"
)

// mysqlExampleRepository struct
type mysqlExampleRepository struct {
	db *sql.DB
}

// NewExampleRepositoryMysql NewUserRepositoryMysql
func NewExampleRepositoryMysql(db *sql.DB) example.Repository {
	return &mysqlExampleRepository{db}
}

// Save Example
func (r *mysqlExampleRepository) Save(me *model.Example) error {

	query := `
	INSERT INTO tbl_users 
	(
		created_at,
		updated_at
	)
	VALUES (?,?,?,?,?,?,?,?)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(me.CreatedAt, me.UpdatedAt)

	if err != nil {
		return err
	}

	lastInsertIdInt64, err := result.LastInsertId()

	if err != nil {
		return err
	}

	lastInsertIdStr := strconv.FormatInt(lastInsertIdInt64, 10)

	me.ExampleID = lastInsertIdStr

	return nil
}

// FindByID Example
func (r *mysqlExampleRepository) FindByID(id string) (*model.Example, error) {

	query := `
	SELECT 
		created_at,
		updated_at
	FROM tbl_users WHERE user_id = ? && deleted_at is NULL`

	var me model.Example

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&me.CreatedAt, &me.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &me, nil
}

// FindAll Example
func (r *mysqlExampleRepository) FindAll(limit, offset, order string) (model.ExampleList, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM tbl_example
	ORDER BY created_at %s
	LIMIT %s 
	OFFSET %s`, order, limit, offset)

	var mel model.ExampleList

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var me model.Example

		err = rows.Scan(&me.CreatedAt, &me.UpdatedAt)

		if err != nil {
			return nil, err
		}
		mel = append(mel, me)
	}

	return mel, nil
}

// Update Example
func (r *mysqlExampleRepository) Update(id string, u *model.Example) (rowAffected *string, err error) {

	query := `
	UPDATE tbl_example 
	SET 
		username=?, 
		email=?, 
	WHERE user_id=?`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(u.CreatedAt, u.UpdatedAt, id)

	if err != nil {
		return nil, err
	}

	rowsAffectedInt64, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	rowsAffectedStr := strconv.FormatInt(rowsAffectedInt64, 10)

	rowAffected = &rowsAffectedStr

	return rowAffected, nil

}

// Delete Example
func (r *mysqlExampleRepository) Delete(id string) error {

	query := `
	UPDATE tbl_users 
	SET	deleted_at = ?
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

// IsExistsByID Example
func (r *mysqlExampleRepository) IsExistsByID(id string) (isExist bool, err error) {

	query := "SELECT EXISTS(SELECT TRUE from example WHERE example_id = ?)"

	statement, err := r.db.Prepare(query)

	if err != nil {
		return false, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&isExist)

	if err != nil {
		return false, err
	}

	return isExist, nil
}
