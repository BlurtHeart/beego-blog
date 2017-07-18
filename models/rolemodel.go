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

func FindRoleById(id int) Role {
	o := orm.NewOrm()
	var role Role
	o.QueryTable(role).Filter("Id", id).One(&role)
	return role
}

func SaveRole(role *Role) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(role)
	return id
}

func UpdateRole(role *Role) {
	o := orm.NewOrm()
	o.Update(role, "RoleName", "Permissions")
}

func DeleteRole(role *Role) {
	o := orm.NewOrm()
	o.Delete(role)
}

func DeleteRoleById(id int) {
	o := orm.NewOrm()
	o.Raw("delete from roles where id=?", id).Exec()
}

func FindRoles() []*Role {
	o := orm.NewOrm()
	var role Role
	var roles []*Role
	o.QueryTable(role).All(&roles)
	return roles
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
