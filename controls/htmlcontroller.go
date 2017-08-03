package controls

import (
	//	"encoding/json"

	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (this *DefaultController) Prepare() {
	param := this.GetSession("username")
	var username string
	if param != nil {
		username = param.(string)
		this.Data["username"] = username
		this.Data["IsAdmin"] = true
	} else {
		this.Data["IsAdmin"] = false
	}
}

func (c *DefaultController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "empty.html"
}

type UserController struct {
	beego.Controller
}

func (this *UserController) Prepare() {
	param := this.GetSession("username")
	var username string
	if param != nil {
		username = param.(string)
		this.Data["username"] = username
		this.Data["IsAdmin"] = true
	} else {
		this.Data["IsAdmin"] = false
	}
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

type PostController struct {
	beego.Controller
}

func (this *PostController) Prepare() {
	param := this.GetSession("username")
	var username string
	if param != nil {
		username = param.(string)
		this.Data["username"] = username
		this.Data["IsAdmin"] = true
	} else {
		this.Data["IsAdmin"] = false
	}
}
func (p *PostController) Get() {
	p.Layout = "layout.tpl"
	p.TplName = "post.html"
}
