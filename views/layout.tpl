<!doctype html>
<html lang="en">

<head>
<link rel="stylesheet" type="text/css" href="/static/bootstrap/css/bootstrap.min.css">
</head>

<body>
<div class="navbar navbar-inverse" role="navigation">
    <div class="container">
        <div navbar-header>
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="    .navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="/about">博客</a>
        </div>
        <div class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
                <li><a href="/index">Home</a></li>
                <li><a href="/about">About</a></li>
            </ul>
            <ul class="nav navbar-nav navbar-right">
                <li><a href="/test">测试</a></li>
                {{ if .IsAdmin }}
                    <li><a href="/moderate">管理评论</a></li>
                    <li><a href="/post">写博客</a></li>
                    <li class="dropdown">                         
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown">用户</a>
                        <ul class="dropdown-menu">
                            <li><a href="/profile">用户资料</a></li>
                            <li><a href="/logout">登出</a></li>
                        </ul>
                    </li>
                {{ else }}
                    <li><a href="#mylogin" data-toggle="modal">登录</a></li>
                    <div class="modal fade" id="mylogin" role="dialog" aria-hidden="true" tabindex="-1" aria-labelledby="myModalLabel" data-backdrop="static" data-keyboard="false">
                    <div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <button type="button" class="close"
                        data-dismiss="modal" aria-hidden="true">
                            &times;
                            </button>
                            <h4 class="modal-title" id="mymodaltitle">用户登录</h4>
                        </div>
                        <div class="loginmessage"></div>
                        <div class="modal-body">
                            <form id="loginform" role="form" class="form-horizontal" action="" method="post">
                                <div class="form-group">
                                    <label class="control-label col-sm-3">用户名</label>
                                    <div class="col-sm-9">
                            <input type="text" class="form-control" placeholder="please input username" name="username"/><br/>
                                    </div>
                                    <label class="control-label col-sm-3">密码</label>
                                    <div class="col-sm-9">
                            <input type="password" class="form-control" placeholder="please input passwd" name="password"/><br/>
                            <input type="submit" class="btn btn-primary" align="right" value="登录"            />
                            <button class="btn btn-danger" data-dismiss="modal">关闭</button>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                    </div>
                    </div>
                    <li><a href="#myregister" data-toggle="modal">注册</a></li>
                    <div class="modal fade" id="myregister" role="dialog" aria-hidden="true" tabindex="-1" aria-labelledby="myModalLabel">
                    <div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <button type="button" class="close"
                        data-dismiss="modal" aria-hidden="true">
                            &times;
                            </button>
                            <h4 class="modal-title" id="mymodaltitle">用户注册</h4>
							<div class="registermessage"></div>
                        </div>
                        <div class="modal-body">
                            <form role="form" id="registerform" onClick="return validate_form(this)" class="form-horizontal" action="" method="post">
                                <div class="form-group">
                                    <label class="control-label col-sm-3">邮箱</label>
                                    <div class="col-sm-9">
                            <input type="text" class="form-control" placeholder="please input your email" name="email"/><br/>
                                    </div>
									<label class="control-label col-sm-3">用户名</label>
                                    <div class="col-sm-9">
                            <input type="text" class="form-control" placeholder="please input your username" name="username"/><br/>
                                    </div>
                                    <label class="control-label col-sm-3">密码</label>
                                    <div class="col-sm-9">
                            <input type="password" class="form-control password1" placeholder="please input your password" name="password"/><br/>
                                    </div>
                                    <label class="control-label col-sm-3">确认密码</label>
                                    <div class="col-sm-9">
                            <input type="password" class="form-control password2" placeholder="please check your password" name="checkpassword"/><br/>
                            <input type="submit" class="btn btn-primary" align="right" value="注册"/>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                    </div>
                    </div>
                {{ end }}
            </ul>
        </div>
    </div>
</div>

<div class="container">
    {{.LayoutContent}}
</div>


<div style="height:60px;"></div>

<nav class="navbar navbar-inverse navbar-fixed-bottom">
    <div class="container"> 
        <p class="navbar-text">Copyright © Steve Cloud</p>
    </div>
</nav>

<script src="/static/js/jquery.js"></script>
<script src="/static/js/bootstrap.js"></script>
<script src="/static/js/base.js"></script>

</body>
