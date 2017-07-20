package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id         int        `orm:"pk;auto"`
	Title      string     `orm:"type(text)"`
	Body       string     `orm:"type(text)"`
	User       *User      `orm:"rel(fk)"`
	Comments   []*Comment `orm:"reverse(many)"`
	CreateTime time.Time  `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time  `orm:"auto_now;type(datetime)"`
}

func (p *Post) TableName() string {
	return "posts"
}

func SavePost(post *Post) int {
	o := orm.NewOrm()
	id, _ := o.Insert(post)
	return int(id)
}

func FindPostById(id int) Post {
	o := orm.NewOrm()
	var post Post
	o.QueryTable(post).Filter("Id", id).One(&post)
	return post
}
