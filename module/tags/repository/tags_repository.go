package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/orgmatileg/go-blog-rest-api-example/module/tags"
	"github.com/orgmatileg/go-blog-rest-api-example/module/tags/model"
)

// mysqlTagsRepository struct
type mysqlTagsRepository struct {
	db *sql.DB
}

// NewTagsRepositoryMysql NewUserRepositoryMysql
func NewTagsRepositoryMysql(db *sql.DB) tags.Repository {
	return &mysqlTagsRepository{db}
}

// Save Exampleexample
func (r *mysqlTagsRepository) Save(me *model.Tag) error {

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

	_, err = statement.Exec(me.TagName)

	if err != nil {
		return err
	}

	return nil
}

// FindAll Example
func (r *mysqlTagsRepository) FindAll(limit, offset string) (model.Tags, error) {

	query := fmt.Sprintf(`
	SELECT DISTINCT tag_name 
	FROM trx_posts_tags
	LIMIT %s 
	OFFSET %s`, limit, offset)

	var lmt model.Tags

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var mt model.Tag

		err = rows.Scan(&mt.TagName)

		if err != nil {
			return nil, err
		}
		lmt = append(lmt, mt)
	}

	return lmt, nil
}

// Delete Example
func (r *mysqlTagsRepository) Delete(tagName string) error {

	query := `
	DELETE FROM trx_posts_tags
	WHERE tag_name = ?`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(time.Now(), tagName)

	if err != nil {
		return err
	}

	return nil
}
