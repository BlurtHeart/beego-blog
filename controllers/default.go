package controllers

import (
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/session"
)

//var globalSessions *session.Manager
//var globalSessionConfig *session.ManagerConfig

func init() {
	//	globalSessionConfig = &session.ManagerConfig{
	//		CookieName:      "gosessionid",
	//		EnableSetCookie: true,
	//		Gclifetime:      3600,
	//		Maxlifetime:     3600,
	//		CookieLifeTime:  3600,
	//		Secure:          false}
	//	globalSessions, _ = session.NewManager("memory", globalSessionConfig)
	//	go globalSessions.GC()
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.Redirect(302, "/index")
}
