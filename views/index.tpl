<!DOCTYPE html>

{{range .posts}}
  <h3><b><a href="/post/{{.Id}}">{{.Title}}</h3></a>
  by:&nbsp;<a href="/profile/{{.User.Id}}">{{.User.Username}}</a></b>
  <h5>{{.Body}}</h5>
{{end}}