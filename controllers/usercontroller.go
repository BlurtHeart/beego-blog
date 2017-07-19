package controllers

import (
	"beego-blog/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("user info")
}

func (c *UserController) Register() {
	var ur models.UserRegisterRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &ur)
	var u models.User
	u.Username = ur.Username
	u.Email = ur.Email
	u.Password_hash = ur.Password
	role := models.FindRoleByName("User")
	u.Roles = append(u.Roles, &role)
	id := models.SaveUser(&u)
	models.SaveUserRole(id, role.Id)
	res := struct {
		Username string
		Id       int
		Email    string
	}{u.Username, id, u.Email}
	c.Data["json"] = &res
	c.ServeJSON()
}

func (c *UserController) Login() {
	var ul models.UserLoginRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &ul)
	u := models.FindUserByName(ul.Username)
	// 0 - username uncorrect
	// 1 - login ok
	// 2 - password uncorrect
	var result int
	if u.Id != 0 {
		if u.Password_hash != ul.Password {
			result = 2
		} else {
			result = 1
		}
	} else {
		result = 0
	}
	res := struct {
		Username string
		Email    string
		Result   int
	}{ul.Username, ul.Email, result}
	c.Data["json"] = &res
	c.ServeJSON()
}
