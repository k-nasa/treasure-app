package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/service"
)

type Comment struct {
	dbx *sqlx.DB
}

func NewComment(dbx *sqlx.DB) *Comment {
	return &Comment{dbx: dbx}
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newComment := &model.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newComment.UserID = &user.ID

	commentService := service.NewCommentService(c.dbx)

	newComment, err = commentService.Create(newComment)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, newComment, nil
}

func (c *Comment) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	cid, err := strconv.ParseInt(id, 10, 64)

	requestComment := &model.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&requestComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	requestComment.ID = cid

	commentService := service.NewCommentService(c.dbx)
	comment, err := commentService.Update(requestComment)

	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, comment, nil
}

// func (c *Comment) Delete(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// }
