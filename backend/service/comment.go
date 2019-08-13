package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Comment struct {
	db *sqlx.DB
}

func NewCommentService(db *sqlx.DB) *Comment {
	return &Comment{db}
}

func (c *Comment) Create(newComment *model.Comment) (*model.Comment, error) {
	comment := &model.Comment{}

	if err := dbutil.TXHandler(c.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateComment(tx, newComment)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		comment, err = repository.FindComment(c.db, id)

		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed comment insert transaction")
	}

	return comment, nil
}

func (c *Comment) Update(updateComment *model.Comment) (*model.Comment, error) {
	_, err := repository.FindComment(c.db, updateComment.ID)

	if err != nil {
		return nil, errors.Wrap(err, "failed find comment")
	}

	comment := &model.Comment{}
	if err := dbutil.TXHandler(c.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateComment(tx, updateComment)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}

		comment, err = repository.FindComment(c.db, updateComment.ID)

		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed comment update transaction")
	}

	return comment, nil
}
