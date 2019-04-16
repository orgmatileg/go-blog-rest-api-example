package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/orgmatileg/go-blog-rest-api-example/module/posts"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"
)

// mysqlPostsRepository struct
type mysqlPostsRepository struct {
	db *sql.DB
}

// NewPostsRepositoryMysql
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
		tx.Rollback()
		return err
	}

	statement, err := tx.Prepare(queryPost)

	if err != nil {
		tx.Rollback()
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(mp.PostImage, mp.PostSubject, mp.PostContent, mp.Author.AuthorID, mp.CreatedAt, mp.UpdatedAt)

	if err != nil {
		tx.Rollback()
		return err
	}

	lastInsertIdInt64, err := result.LastInsertId()

	if err != nil {
		return err
	}

	lastInsertIdStr := strconv.FormatInt(lastInsertIdInt64, 10)

	mp.PostID = lastInsertIdStr

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

		_, err = statement.Exec(lastInsertIdInt64, "No Tag")

		if err != nil {
			tx.Rollback()
			return err
		}

	} else {
		queryTags = `
		INSERT INTO trx_posts_tags (post_id, tag_name)
		VALUES
		`

		valuesTags := strings.Repeat(fmt.Sprintf(",(%d,?)", lastInsertIdInt64), len(mp.Tags))

		statement, err = tx.Prepare(queryTags + valuesTags[1:])

		if err != nil {
			tx.Rollback()
			return err
		}

		var iTags []interface{} = make([]interface{}, len(mp.Tags))
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

// // FindAll Example
// func (r *mysqlExampleRepository) FindAll(limit, offset, order string) (model.ExampleList, error) {

// 	query := fmt.Sprintf(`
// 	SELECT *
// 	FROM tbl_example
// 	ORDER BY created_at %s
// 	LIMIT %s
// 	OFFSET %s`, order, limit, offset)

// 	var mel model.ExampleList

// 	rows, err := r.db.Query(query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var me model.Example

// 		err = rows.Scan(&me.CreatedAt, &me.UpdatedAt)

// 		if err != nil {
// 			return nil, err
// 		}
// 		mel = append(mel, me)
// 	}

// 	return mel, nil
// }

// // Update Example
// func (r *mysqlExampleRepository) Update(id string, u *model.Example) (rowAffected *string, err error) {

// 	query := `
// 	UPDATE tbl_example
// 	SET
// 		username=?,
// 		email=?,
// 	WHERE user_id=?`

// 	statement, err := r.db.Prepare(query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer statement.Close()

// 	result, err := statement.Exec(u.CreatedAt, u.UpdatedAt, id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	rowsAffectedInt64, err := result.RowsAffected()

// 	if err != nil {
// 		return nil, err
// 	}

// 	rowsAffectedStr := strconv.FormatInt(rowsAffectedInt64, 10)

// 	rowAffected = &rowsAffectedStr

// 	return rowAffected, nil

// }

// // Delete Example
// func (r *mysqlExampleRepository) Delete(id string) error {

// 	query := `
// 	UPDATE tbl_users
// 	SET	deleted_at = ?
// 	WHERE user_id = ?`

// 	statement, err := r.db.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer statement.Close()

// 	_, err = statement.Exec(time.Now(), id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // IsExistsByID Example
// func (r *mysqlExampleRepository) IsExistsByID(id string) (isExist bool, err error) {

// 	query := "SELECT EXISTS(SELECT TRUE from example WHERE example_id = ?)"

// 	statement, err := r.db.Prepare(query)

// 	if err != nil {
// 		return false, err
// 	}

// 	defer statement.Close()

// 	err = statement.QueryRow(id).Scan(&isExist)

// 	if err != nil {
// 		return false, err
// 	}

// 	return isExist, nil
// }
