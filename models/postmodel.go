package models

import (
	"time"
)

type Post struct {
	Id         int        `orm:"pk;auto"`
	Title      string     `orm:"type(text)"`
	Body       string     `orm:"type(text)"`
	BodyHtml   string     `orm:"type(text)"`
	User       *User      `orm:"rel(fk)"`
	Comments   []*Comment `orm:"reverse(many)"`
	CreateTime time.Time  `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time  `orm:"auto_now;type(datetime)"`
}

func (p *Post) TableName() string {
	return "posts"
}
