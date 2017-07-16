package models

type Comment struct {
	Id         int    `orm:"pk;auto"`
	Body       string `orm:"type(text)"`
	BodyHtml   string `orm:"type(text)"`
	User       *User  `orm:"rel(fk)"`
	Post       *post  `orm:"rel(fk)"`
	Disabled   bool
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Comment) TableName() string {
	return "comments"
}
