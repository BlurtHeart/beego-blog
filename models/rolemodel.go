package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Role struct {
	Id          int     `orm:"pk;auto"`
	RoleName    string  `orm:"size(20);unique"`
	Users       []*User `orm:"reverse(many)"`
	Permissions int
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func (r *Role) TableName() string {
	return "roles"
}

func FindRoleByName(rolename string) Role {
	o := orm.NewOrm()
	var role Role
	o.QueryTable(role).Filter("RoleName", rolename).One(&role)
	return role
}

func InsertRoles() {
	roles := []Role{
		Role{RoleName: "User", Permissions: FOLLOW | COMMENT | WRITE_ARTICLES},
		Role{RoleName: "Moderator", Permissions: FOLLOW | COMMENT | WRITE_ARTICLES | MODERATE_COMMENTS},
		Role{RoleName: "Administrator", Permissions: 0xff},
	}
	o := orm.NewOrm()
	for _, r := range roles {
		qr := FindRoleByName(r.RoleName)
		if qr.Id == 0 {
			o.Insert(&r)
		}
	}
}

func init() {
	orm.RegisterModel(new(Role), new(User), new(Comment), new(Post))
}
