package controls

import (
	"strconv"
	//	"encoding/json"
	"beego-blog/models"

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
	c.TplName = "index.tpl"
	ps := models.FindAllPosts()
	c.Data["posts"] = ps
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
		this.Data["username"] = ""
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

func (u *UserController) Logout() {
	// u.DelSession(u.Data["username"].(string))
	u.DestroySession()
	u.Redirect("/", 302)
}

func (u *UserController) Profile() {
	u.Layout = "layout.tpl"
	u.TplName = "profile.html"
	user := models.FindUserByName(u.Data["username"].(string))
	u.Data["user"] = user
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

func (this *PostController) Detail() {
	post_id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	post := models.FindPostById(post_id)
	this.Data["post"] = post
	this.Layout = "layout.tpl"
	this.TplName = "post-detail.html"
}
