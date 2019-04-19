package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/orgmatileg/go-blog-rest-api-example/module/posts"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"
)

// mysqlPostsRepository struct
type mysqlPostsRepository struct {
	db *sql.DB
}

// NewPostsRepositoryMysql struct
func NewPostsRepositoryMysql(db *sql.DB) posts.Repository {
	return &mysqlPostsRepository{db}
}

// Save Example
func (r *mysqlPostsRepository) Save(mp *model.Post) error {

	// Post
	queryPost := `
	INSERT INTO tbl_posts 
	(
		post_featured_image,
		post_subject,
		post_content,
		author,
		created_at,
		updated_at
	)
	VALUES (?, ?, ?, ?, ?, ?)`

	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	statement, err := tx.Prepare(queryPost)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(mp.PostImage, mp.PostSubject, mp.PostContent, mp.Author.AuthorID, mp.CreatedAt, mp.UpdatedAt)

	if err != nil {
		tx.Rollback()
		return err
	}

	lastInsertIDInt64, err := result.LastInsertId()

	if err != nil {
		return err
	}

	lastInsertIDStr := strconv.FormatInt(lastInsertIDInt64, 10)

	mp.PostID = lastInsertIDStr

	// Tags
	var queryTags string

	if mp.Tags == nil {

		queryTags = `
		INSERT INTO trx_posts_tags (post_id, tag_name)
		VALUES (?, ?)
		`

		statement, err = tx.Prepare(queryTags)

		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = statement.Exec(lastInsertIDInt64, "No Tag")

		if err != nil {
			tx.Rollback()
			return err
		}

	} else {
		queryTags = `
		INSERT INTO trx_posts_tags (post_id, tag_name)
		VALUES
		`

		valuesTags := strings.Repeat(fmt.Sprintf(",(%d,?)", lastInsertIDInt64), len(mp.Tags))

		statement, err = tx.Prepare(queryTags + valuesTags[1:])

		if err != nil {
			tx.Rollback()
			return err
		}

		iTags := make([]interface{}, len(mp.Tags))
		for i, tag := range mp.Tags {
			iTags[i] = tag
		}

		_, err = statement.Exec(iTags...)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

// FindByID Example
func (r *mysqlPostsRepository) FindByID(id string) (*model.Post, error) {

	queryPost := `
	SELECT * 
	FROM tbl_v_posts
	WHERE post_id = ? and is_publish = 1`

	var mp model.Post

	statement, err := r.db.Prepare(queryPost)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&mp.PostID, &mp.PostImage, &mp.PostSubject, &mp.PostContent, &mp.Author.AuthorID, &mp.IsPublish, &mp.CreatedAt, &mp.UpdatedAt, &mp.Author.AuthorFullName, &mp.Author.AuthorPhotoProfile)

	if err != nil {
		return nil, err
	}

	queryTags := `
	SELECT tag_name
	FROM trx_posts_tags
	WHERE post_id = ?`

	statement, err = r.db.Prepare(queryTags)

	if err != nil {
		return nil, err
	}

	row, err := statement.Query(id)

	if err != nil {
		return nil, err
	}

	var tags []string

	for row.Next() {
		var tag string

		err = row.Scan(&tag)

		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	mp.Tags = tags

	return &mp, nil
}

// FindAll Example
func (r *mysqlPostsRepository) FindAll(limit, offset, order string) (model.Posts, error) {

	queryPost := fmt.Sprintf(`
	SELECT *
	FROM tbl_v_posts
	WHERE is_publish = 1
	ORDER BY created_at %s
	LIMIT %s
	OFFSET %s`, order, limit, offset)

	var lmp model.Posts

	rows, err := r.db.Query(queryPost)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var mp model.Post

		err = rows.Scan(&mp.PostID, &mp.PostImage, &mp.PostSubject, &mp.PostContent, &mp.Author.AuthorID, &mp.IsPublish, &mp.CreatedAt, &mp.UpdatedAt, &mp.Author.AuthorFullName, &mp.Author.AuthorPhotoProfile)

		if err != nil {
			return nil, err
		}

		queryTags := `
		SELECT tag_name
		FROM trx_posts_tags
		WHERE post_id = ?`

		statement, err := r.db.Prepare(queryTags)

		if err != nil {
			return nil, err
		}

		row, err := statement.Query(mp.PostID)

		if err != nil {
			return nil, err
		}

		var tags []string

		for row.Next() {
			var tag string

			err = row.Scan(&tag)

			if err != nil {
				return nil, err
			}

			tags = append(tags, tag)
		}

		mp.Tags = tags

		lmp = append(lmp, mp)
	}

	return lmp, nil
}

// Update Example
func (r *mysqlPostsRepository) Update(id string, mp *model.Post) (rowAffected *string, err error) {

	queryPost := `
	UPDATE tbl_posts
	SET
		post_featured_image = ?,
		post_subject = ?,
		post_content = ?,
		is_publish = ?,
		updated_at = ?
	WHERE post_id = ?`

	statement, err := r.db.Prepare(queryPost)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(mp.PostImage, mp.PostSubject, mp.PostContent, mp.IsPublish, time.Now(), id)

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
func (r *mysqlPostsRepository) Delete(id string) error {

	query := `
	DELETE FROM tbl_posts
	WHERE post_id = ?`

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

// IsExistsByID Example
func (r *mysqlPostsRepository) IsExistsByID(id string) (isExist bool, err error) {

	query := "SELECT EXISTS(SELECT TRUE from tbl_posts WHERE post_id = ?)"

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
