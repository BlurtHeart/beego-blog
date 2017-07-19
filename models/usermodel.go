package models

import (
	"time"

	"github.com/astaxie/beego/orm"
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

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id            int        `orm:"pk;auto"`
	Username      string     `orm:"unique;index"`
	Email         string     `orm:"size(128);index"`
	Password_hash string     `orm:"size(128)"`
	Confirmed     bool       `orm:"default(false)"`
	Aboutme       string     `orm:"null;type(text)"`
	Roles         []*Role    `orm:"rel(m2m)"`
	Posts         []*Post    `orm:"reverse(many)"`
	Comments      []*Comment `orm:"reverse(many)"`
	CreateTime    time.Time  `orm:"auto_now_add;type(datetime)"`
	UpdateTime    time.Time  `orm:"auto_now;type(datetime)"`
}

func (u *User) TableName() string {
	return "users"
}

func FindUserByName(username string) User {
	o := orm.NewOrm()
	var user User
	o.QueryTable(user).Filter("Username", username).One(&user)
	o.Read(&user)
	o.LoadRelated(&user, "Roles")
	return user
}

func FindUserById(id int) User {
	o := orm.NewOrm()
	var user User
	o.QueryTable(user).Filter("Id", id).One(&user)
	o.Read(&user)
	o.LoadRelated(&user, "Roles")
	return user
}

func FindUsers() []*User {
	o := orm.NewOrm()
	var user User
	var users []*User
	o.QueryTable(user).All(&users)
	return users
}

func SaveUser(user *User) int {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return int(id)
}

func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

func DeleteUser(user *User) {
	o := orm.NewOrm()
	o.Delete(user)
}

func DeleteUserById(id int) {
	o := orm.NewOrm()
	o.Raw("delete from roles where id=?", id).Exec()
}

func SaveUserRole(user_id, role_id int) {
	o := orm.NewOrm()
	o.Raw("insert into users_roless(users_id, roles_id)values(?, ?)", user_id, role_id).Exec()
}
