package controllers

import (
	"beego-blog/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Post() {
	type commentResponse struct {
		CommentId int    `json:"comment_id"`
		Result    int    `json:"result"`
		Message   string `json:"message"`
		Next      string `json:"next"`
	}

	param := c.GetSession("user_id")
	var user_id int
	if param != nil {
		user_id = param.(int)
	} else {
		res := commentResponse{Result: 0, Message: "need login", Next: "/login"}
		c.Data["json"] = &res
		c.ServeJSON()
		return
	}

	var message string
	var next string

	type commentRequest struct {
		Body   string `json:"body"`
		PostId string `json:"post_id"`
	}
	var cr commentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &cr)
	post_id, _ := strconv.Atoi(cr.PostId)
	u := models.FindUserById(user_id)
	post := models.FindPostById(post_id)

	var comment models.Comment
	comment.Body = cr.Body
	comment.User = &u
	comment.Post = &post

	var result int
	id, excResult := models.SaveComment(&comment)
	if excResult {
		result = 1
		message = "comment ok"
	} else {
		result = 1
		message = "comment failed"
	}

	res := commentResponse{id, result, message, next}
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
