package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Article struct {
	db *sqlx.DB
}

func NewArticle(db *sqlx.DB) *Article {
	return &Article{db}
}

func (a *Article) GetArticleComment(id int64) (*model.ArticleResp, error) {
	article, err := repository.FindArticle(a.db, id)

	if err != nil {
		return nil, err
	}

	comments, err := repository.FindCommentsByAirticleID(a.db, article.ID)

	if err != nil && err == sql.ErrNoRows {
		return &model.ArticleResp{article, make([]*model.Comment, 0)}, nil
	} else if err != nil {
		return nil, err
	}

	return &model.ArticleResp{article, comments}, nil
}

func (a *Article) Update(id int64, newArticle *model.Article) error {
	_, err := repository.FindArticle(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateArticle(tx, id, newArticle)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed article update transaction")
	}
	return nil
}

func (a *Article) Destroy(id int64) error {
	_, err := repository.FindArticle(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyArticle(tx, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed article delete transaction")
	}
	return nil
}

func (a *Article) Create(newArticle *model.ArticleTag) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateArticle(tx, newArticle.Article)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdId = id

		for _, tag_id := range newArticle.TagIDs {
			repository.CreateTag(tx, createdId, tag_id)
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed article insert transaction")
	}
	return createdId, nil
}
