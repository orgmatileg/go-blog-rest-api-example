package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/orgmatileg/go-blog-rest-api-example/module/contact_us"
	"github.com/orgmatileg/go-blog-rest-api-example/module/contact_us/model"
)

// mysqlContactUsRepository struct
type mysqlContactUsRepository struct {
	db *sql.DB
}

// NewContactUsRepositoryMysql
func NewContactUsRepositoryMysql(db *sql.DB) contact_us.Repository {
	return &mysqlContactUsRepository{db}
}

// Save Contact Us
func (r *mysqlContactUsRepository) Save(c *model.ContactUs) error {

	query := `INSERT INTO tbl_contact_us (
	full_name,
	email,
	subject,
	message,
	created_at
	)
	VALUES (?,?,?,?,?)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(c.FullName, c.Email, c.Subject, c.Message, c.CreatedAt)

	if err != nil {
		return err
	}

	lastInsertIdInt64, err := result.LastInsertId()

	if err != nil {
		return err
	}

	lastInsertIdStr := strconv.FormatInt(lastInsertIdInt64, 10)

	c.ContactUsID = lastInsertIdStr

	return nil
}

// FindByID User
func (r *mysqlContactUsRepository) FindByID(id string) (*model.ContactUs, error) {

	query := `
	SELECT *
	FROM tbl_contact_us 
	WHERE contact_us_id = ?`

	var mc model.ContactUs

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&mc.ContactUsID, &mc.FullName, &mc.Email, &mc.Subject, &mc.Message, &mc.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &mc, nil
}

// FindAll User
func (r *mysqlContactUsRepository) FindAll(limit, offset, order string) (model.ContactUsList, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM tbl_contact_us
	ORDER BY created_at %s
	LIMIT %s
	OFFSET %s`, order, limit, offset)

	var mcl model.ContactUsList

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var mc model.ContactUs

		err = rows.Scan(&mc.ContactUsID, &mc.FullName, &mc.Email, &mc.Subject, &mc.Message, &mc.CreatedAt)

		if err != nil {
			return nil, err
		}
		mcl = append(mcl, mc)
	}

	return mcl, nil
}

// Delete User
func (r *mysqlContactUsRepository) Delete(id string) error {

	query := `
	DELETE FROM tbl_contact_us
	WHERE contact_us_id = ?`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

// Count Posts
func (r *mysqlContactUsRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM tbl_contact_us
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return -1, err
	}

	defer statement.Close()

	err = statement.QueryRow().Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}
