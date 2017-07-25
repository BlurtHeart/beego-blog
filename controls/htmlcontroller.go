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

func (u *UserController) Login() {
	u.Layout = "layout.tpl"
	u.TplName = "login.html"
}
