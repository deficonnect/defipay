package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
    globalSessions, _ = session.NewManager("memory", &session.ManagerConfig{
		CookieName: "deficonnect",
		EnableSetCookie: true,
		Gclifetime: 3600,
		Maxlifetime: 3600,
		Secure: true,
		CookieLifeTime: 3600,
	})
    go globalSessions.GC()
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	c.Controller.Prepare()
	c.Layout = "layout.html"
	c.LayoutSections = map[string]string{}
}

func (c *MainController) Get() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
    defer sess.SessionRelease(c.Ctx.ResponseWriter)
		
	c.Data["Debug"] = c.GetString("debug")
	c.Data["JsController"] = "index"
	c.LayoutSections["Scripts"] = "index_scripts.html"
	c.TplName = "index.html"
}

func (c *MainController) Form() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
    defer sess.SessionRelease(c.Ctx.ResponseWriter)
	
	c.Data["Debug"] = c.GetString("debug")
	c.Data["JsController"] = "form"
	c.TplName = "form.html"
}
