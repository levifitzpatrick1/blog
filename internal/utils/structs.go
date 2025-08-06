package utils

import (
	"html/template"
	"time"
)

type Tag struct {
	Name  string
	Color string
}

type TempPost struct {
	Title      string    `yaml:"title"`
	Slug       string    `yaml:"slug"`
	Blurb      string    `yaml:"blurb"`
	CreateDate time.Time `yaml:"date"`
	ModifyDate time.Time `yaml:"modifieddate"`
	Tags       []string  `yaml:"tags"`
}

type Post struct {
	Title       string
	Slug        string
	Blurb       string
	CreateDate  time.Time
	ModifyDate  time.Time
	Tags        []Tag
	Content     string
	HTMLContent template.HTML
}

func (p *Post) GetTags(tagNames []string) {
	p.Tags = make([]Tag, len(tagNames))
	for i, name := range tagNames {
		color, ok := TagColorMap[name]
		if !ok {
			color = ColorDefault
		}
		p.Tags[i] = Tag{Name: name, Color: color}
	}
}
