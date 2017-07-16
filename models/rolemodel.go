package models

import (
	"time"
)

type Role struct {
	Id          int     `orm:"pk;auto"`
	RoleName    string  `orm:"size(20);unique"`
	Users       []*User `orm:"reverse(many)"`
	Permissions int
}

func (r *Role) TableName() string {
	return "roles"
}
