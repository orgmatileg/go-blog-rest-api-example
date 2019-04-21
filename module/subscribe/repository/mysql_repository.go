package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/orgmatileg/go-blog-rest-api-example/module/subscribe"
	"github.com/orgmatileg/go-blog-rest-api-example/module/subscribe/model"
)

// mysqlExampleRepository struct
type mysqlSubscribeRepository struct {
	db *sql.DB
}

// NewSubscribeRepositoryMysql NewUserRepositoryMysql
func NewSubscribeRepositoryMysql(db *sql.DB) subscribe.Repository {
	return &mysqlSubscribeRepository{db}
}

// Save Example
func (r *mysqlSubscribeRepository) Save(ms *model.Subscribe) error {

	query := `
	INSERT INTO tbl_subscribe 
	(
		email,
		created_at
	)
	VALUES (?, ?)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(ms.Email, ms.CreatedAt)

	if err != nil {
		return err
	}

	lastInsertIdInt64, err := result.LastInsertId()

	if err != nil {
		return err
	}

	lastInsertIdStr := strconv.FormatInt(lastInsertIdInt64, 10)

	ms.SubscribeID = lastInsertIdStr

	return nil
}

// FindByID Example
func (r *mysqlSubscribeRepository) FindByID(id string) (*model.Subscribe, error) {

	query := `
	SELECT *
	FROM tbl_subscribe WHERE subscribe_id = ?`

	var ms model.Subscribe

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&ms.SubscribeID, &ms.Email, &ms.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &ms, nil
}

// FindAll Example
func (r *mysqlSubscribeRepository) FindAll(limit, offset, order string) (model.SubscribeList, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM tbl_subscribe
	ORDER BY created_at %s
	LIMIT %s 
	OFFSET %s`, order, limit, offset)

	var msl model.SubscribeList

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ms model.Subscribe

		err = rows.Scan(&ms.SubscribeID, &ms.Email, &ms.CreatedAt)

		if err != nil {
			return nil, err
		}
		msl = append(msl, ms)
	}

	return msl, nil
}
