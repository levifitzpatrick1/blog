package utils

import db "github.com/levifitzpatrick1/blog/internal/database/generated"

type BlogItem struct {
	Blog db.Blog
	Tags []db.Tag
}

type BlogsData struct {
	Posts       []BlogItem
	CurrentSort string
	CurrentPath string
}
