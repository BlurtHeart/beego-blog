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

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	var ul models.UserLoginRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &ul)
	// var res models.UserLoginResponse
	// res.Username = ul.Username
	// res.Result = true
	res := struct {
		Username string
		Result   bool
	}{ul.Username, true}
	c.Data["json"] = &res
	c.ServeJSON()
}
