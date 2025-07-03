package utils

import (
	"context"

	db "github.com/levifitzpatrick1/blog/internal/database/generated"
)

func GetTagsBlog(id int64, q *db.Queries) ([]db.Tag, error) {
	tags, err := q.GetTagsBlog(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func GetTagsFrc(id int64, q *db.Queries) ([]db.Tag, error) {
	tags, err := q.GetTagsFrc(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
