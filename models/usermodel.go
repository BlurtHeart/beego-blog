package models

import (
	"time"
)

type UserLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Result   bool   `json:"result"`
}

type User struct {
	Id            int        `orm:"pk;auto"`
	Username      string     `orm:"unique;index"`
	Email         string     `orm:"size(128);index"`
	Password_hash string     `orm:"size(128)"`
	Confirmed     bool       `orm:"default(false)"`
	Aboutme       string     `orm:"null;type(text)"`
	RoleID        []int      `orm:"column(role_id);rel(m2m)"`
	Posts         []*Post    `orm:"reverse(many)"`
	Comments      []*Comment `orm:"reverse(many)"`
	CreateTime    time.Time  `orm:"auto_now_add;type(datetime)"`
	UpdateTime    time.Time  `orm:"auto_now;type(datetime)"`
}

func (u *User) TableName() string {
	return "users"
}
