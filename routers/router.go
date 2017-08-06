package routers

import (
	"beego-blog/controllers"
	"beego-blog/controls"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/user", &controllers.UserController{})

	beego.Router("/api/login", &controllers.UserController{}, "post:Login")
	beego.Router("/api/register", &controllers.UserController{}, "post:Register")
	beego.Router("/api/post", &controllers.PostController{})
	beego.Router("/api/comment", &controllers.CommentController{})

	/**************************************************/
	beego.Router("/register", &controls.UserController{}, "get:Register")
	beego.Router("/login", &controls.UserController{}, "get:Login")
	beego.Router("/logout", &controls.UserController{}, "get:Logout")
	beego.Router("/index", &controls.DefaultController{})
	beego.Router("/post", &controls.PostController{})
	beego.Router("/post/:id([0-9]+", &controls.PostController{}, "get:Detail")
	beego.Router("/profile", &controls.UserController{}, "get:Profile")
}
