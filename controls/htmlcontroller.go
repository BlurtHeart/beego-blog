package controls

import (
	//	"encoding/json"

	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (c *DefaultController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "empty.html"
}

type UserController struct {
	beego.Controller
}

func (u *UserController) Register() {
	u.Data["IsAdmin"] = false
	u.Layout = "layout.tpl"
	u.TplName = "register.html"
}

func (u *UserController) RegisterPost() {
	type RegisterRequest struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}
	var rr RegisterRequest
	//	json.Unmarshal(u.Ctx.Input.RequestBody, &rr)
	rr.UserName = u.GetString("username")
	rr.PassWord = u.GetString("password")
	u.Data["json"] = &rr
	u.ServeJSON()
}

func (u *UserController) Login() {
	u.Layout = "layout.tpl"
	u.TplName = "login.html"
}
