package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateComment(db *sqlx.Tx, c *model.Comment) (sql.Result, error) {
	stmt, err := db.Prepare(`
	insert into comments (user_id, article_id, body) value(?, ?, ?)
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.Exec(c.UserID, c.ArticleID, c.Body)
}

func FindComment(db *sqlx.DB, id int64) (*model.Comment, error) {
	c := model.Comment{}
	if err := db.Get(&c, `
	SELECT id, article_id, user_id, body, created_at, updated_at FROM comments WHERE id = ?
 `, id); err != nil {
		return nil, err
	}

	return &c, nil
}

func FindCommentsByAirticleID(db *sqlx.DB, articleID int64) ([]*model.Comment, error) {
	c := []*model.Comment{}
	if err := db.Select(&c, `
	SELECT id, article_id, user_id, body, created_at, updated_at FROM comments WHERE article_id = ?
 `, articleID); err != nil {
		return nil, err
	}

	return c, nil
}

func UpdateComment(db *sqlx.Tx, c *model.Comment) (sql.Result, error) {
	stmt, err := db.Prepare(`update comments set body = ? where id = ?`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.Exec(c.Body, c.ID)
}
