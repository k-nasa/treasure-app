package model

type Article struct {
	ID     int64  `db:"id" json:"id"`
	UserID *int64 `db:"user_id" json:"user_id"`
	Title  string `db:"title" json:"title"`
	Body   string `db:"body" json:"body"`
}

type ArticleResp struct {
	Article  *Article
	Comments []*Comment
}

type ArticleTag struct {
	Article *Article `json:"article"`
	TagIDs  []int64  `json:"tag_ids"`
}
