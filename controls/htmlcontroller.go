package controls

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Register() {
	u.Data["IsAdmin"] = false
	u.Layout = "layout.tpl"
	u.TplName = "register.html"
}
