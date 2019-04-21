package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/orgmatileg/go-blog-rest-api-example/module/settings"
	"github.com/orgmatileg/go-blog-rest-api-example/module/settings/model"
)

// mysqlExampleRepository struct
type mysqlSettingsRepository struct {
	db *sql.DB
}

// NewExampleRepositoryMysql NewUserRepositoryMysql
func NewSettingsRepositoryMysql(db *sql.DB) settings.Repository {
	return &mysqlSettingsRepository{db}
}

// FindByID Example
func (r *mysqlSettingsRepository) FindByID(id string) (*model.Setting, error) {

	query := `
	SELECT *
	FROM tbl_settings 
	WHERE setting_id = ?
	`

	var ms model.Setting

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&ms.SettingID, &ms.Type, &ms.Name, &ms.FieldName, &ms.Value)

	if err != nil {
		return nil, err
	}

	return &ms, nil
}

// FindAll Example
func (r *mysqlSettingsRepository) FindAll(limit, offset string) (model.Settings, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM tbl_settings
	LIMIT %s 
	OFFSET %s`, limit, offset)

	var msl model.Settings

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ms model.Setting

		err = rows.Scan(&ms.SettingID, &ms.Type, &ms.Name, &ms.FieldName, &ms.Value)

		if err != nil {
			return nil, err
		}
		msl = append(msl, ms)
	}

	return msl, nil
}

// Update Example
func (r *mysqlSettingsRepository) Update(id string, s *model.Setting) (rowAffected *string, err error) {

	query := `
	UPDATE tbl_settings
	SET value = ?  
	WHERE setting_id = ?
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(s.Value, id)

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

// Count Posts
func (r *mysqlSettingsRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM tbl_settings
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
