package controllers

import (
	"beego-blog/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type PostController struct {
	beego.Controller
}

func (p *PostController) Post() {
	param := p.GetSession("username")
	var username string
	if param != nil {
		username = param.(string)
	} else {
		p.Ctx.WriteString("no session found")
	}
	u := models.FindUserByName(username)

	type postrequest struct {
		Body   string
		Title  string
		UserId int `json:"user_id"`
	}
	var pm postrequest
	json.Unmarshal(p.Ctx.Input.RequestBody, &pm)
	// u := models.FindUserById(pm.UserId)

	var pp models.Post
	pp.Body = pm.Body
	pp.Title = pm.Title
	pp.User = &u
	var result int
	var message string
	id, excResult := models.SavePost(&pp)
	if !excResult {
		result = 0
		message = "post failed"
	} else {
		result = 1
		message = "post ok"
	}
	next := fmt.Sprintf("%s%d", "/post/", id)
	res := struct {
		PostId  int    `json:"post_id"` // post id
		Result  int    `json:"result"`  // post result
		Message string `json:"message"`
		Next    string `json:"next"`
	}{id, result, message, next}
	p.Data["json"] = &res
	p.ServeJSON()
}

func (p *PostController) Get() {
	id, _ := p.GetInt("post_id")
	var pp models.Post
	pp = models.FindPostById(id)
	res := struct {
		PostId     int       `json:"post_id"`
		Title      string    `json:"title"`
		Body       string    `json:"body"`
		UserId     int       `json:"user_id"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
	}{pp.Id, pp.Title, pp.Body, pp.User.Id, pp.CreateTime, pp.UpdateTime}
	p.Data["json"] = &res
	p.ServeJSON()
}

func (p *PostController) Delete() {
	id, _ := p.GetInt("post_id")
	post := models.FindPostById(id)
	result := models.DeletePost(&post)

	res := struct {
		PostId int  `json:"post_id"`
		Result bool `json:"result"`
	}{id, result}
	p.Data["json"] = &res
	p.ServeJSON()
}
