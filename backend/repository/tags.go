package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func CreateTag(db *sqlx.Tx, article_id, tag_id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
	insert into article_tag(article_id, tag_id) value(?, ?)
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.Exec(article_id, tag_id)
}
