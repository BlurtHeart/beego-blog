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

	var result int
	var message string
	var id int

	uv := models.FindUserByName(u.Username)
	if uv.Id != 0 {
		result = 2
		message = "用户已被注册"
	} else {
		role := models.FindRoleByName("User")
		u.Roles = append(u.Roles, &role)
		if id, excResult := models.SaveUser(&u); !excResult {
			result = 3
			message = "未知错误"
		} else {
			models.SaveUserRole(id, role.Id)
			result = 1
			message = "注册成功"
		}
	}
	res := struct {
		Username string
		Id       int
		Email    string
		Message  string `json:"message"`
		Result   int    `json:"result"`
	}{u.Username, id, u.Email, message, result}
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
	var message string
	next := "/"
	if u.Id != 0 {
		if u.Password_hash != ul.Password {
			result = 2
			message = "password error"
		} else {
			result = 1
			message = "login success"
			c.SetSession("username", ul.Username)
		}
	} else {
		result = 0
		message = "username error"
	}
	res := struct {
		Username string
		Email    string
		Message  string `json:"message"`
		Next     string `json:"next"`
		Result   int    `json:"result"`
	}{ul.Username, ul.Email, message, next, result}
	c.Data["json"] = &res
	c.ServeJSON()
}
