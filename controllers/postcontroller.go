package controllers

import (
	"beego-blog/models"
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
)

type PostController struct {
	beego.Controller
}

func (p *PostController) Post() {
	type postrequest struct {
		Body   string
		Title  string
		UserId int `json:"user_id"`
	}
	var pm postrequest
	json.Unmarshal(p.Ctx.Input.RequestBody, &pm)
	u := models.FindUserById(pm.UserId)

	var pp models.Post
	pp.Body = pm.Body
	pp.Title = pm.Title
	pp.User = &u
	var result int
	id, excResult := models.SavePost(&pp)
	if !excResult {
		result = 1
	} else {
		result = 0
	}
	res := struct {
		PostId int `json:"post_id"` // post id
		Result int `json:"result"`  // post result
	}{id, result}
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
