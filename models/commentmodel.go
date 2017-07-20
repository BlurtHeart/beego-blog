package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id         int    `orm:"pk;auto"`
	Body       string `orm:"type(text)"`
	User       *User  `orm:"rel(fk)"`
	Post       *Post  `orm:"rel(fk)"`
	Disabled   bool
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Comment) TableName() string {
	return "comments"
}

func SaveComment(comment *Comment) int {
	o := orm.NewOrm()
	id, _ := o.Insert(comment)
	return int(id)
}
