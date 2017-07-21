package controllers

import (
	"beego-blog/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Post() {
	type commentRequest struct {
		Body   string
		UserId int `json:"user_id"`
		PostId int `json:"post_id"`
	}
	var cr commentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &cr)
	u := models.FindUserById(cr.UserId)
	post := models.FindPostById(cr.PostId)

	var comment models.Comment
	comment.Body = cr.Body
	comment.User = &u
	comment.Post = &post
	id := models.SaveComment(&comment)
	fmt.Println(id)

	res := struct {
		CommentId int `json:"comment_id"`
		UserId    int `json:"user_id"`
		PostId    int `json:"post_id"`
		Result    int `json:"result"`
	}{id, cr.UserId, cr.PostId, 1}
	c.Data["json"] = &res
	c.ServeJSON()
}

func (c *CommentController) Get() {
	id, _ := c.GetInt("comment_id")
	comment := models.FindCommentById(id)

	res := struct {
		PostTitle   string `json:"post_title"`
		PostId      int    `json:"post_id"`
		UserId      int    `json:"user_id"`
		UserName    string `json:"username"`
		CommentId   int    `json:"comment_id"`
		CommentBody string `json:"comment_body"`
	}{comment.Post.Title, comment.Post.Id, comment.User.Id, comment.User.Username, comment.Id, comment.Body}
	c.Data["json"] = &res
	c.ServeJSON()
}

func (c *CommentController) Delete() {
	id, _ := c.GetInt("comment_id")
	comment := models.FindCommentById(id)
	result := models.DeleteComment(&comment)

	res := struct {
		CommentId int  `json:"comment_id"`
		Result    bool `json:"result"`
	}{id, result}
	c.Data["json"] = &res
	c.ServeJSON()
}
