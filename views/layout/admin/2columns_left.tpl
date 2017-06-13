<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Mark的私人博客</title>

    <link rel="icon" href="/static/images/favicon.ico" type="image/x-icon"/>
    <link rel="shortcut icon" href="/static/images/favicon.ico" type="image/x-icon"/>
    <link href="/static/css/app.css" rel="stylesheet">
    <link href="/static/css/admin/admin.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <script src="/static/js/jquery-1.9.1.min.js" type="text/javascript"></script>
</head>
<body class="admin_body">
    <nav class="navbar navbar-default navbar-static-top">
        <div class="navbar-header">
            <a class="navbar-brand" href="/"><img src="/static/images/logo.png" /></a>
        </div>
        <ul class="nav navbar-top-links navbar-right">
            <li class="dropdown active">
                <a class="dropdown-toggle active" data-toggle="dropdown" href="###">
                    <i class="fa fa-user fa-fw"></i>  {{.loginUser.Name}}  <i class="fa fa-caret-down"></i>
                </a>
                <ul class="dropdown-menu dropdown-user in">
                    <li><a href="/admin/logout"><i class="fa fa-sign-out fa-fw"></i> Logout</a></li>
                </ul>
            </li>
        </ul>
    </nav>
<div class="main-content">
    <div class="left-menu-aside">
        <div class="menu_dropdown" role="navigation">
            <ul>
                <li {{if is_active .current_url "/admin/article"}}class="active"{{end}}>
                    <a href="/admin/article"><i class="fa fa-pencil-square-o fa-fw"></i> 文章管理</a>
                </li>
                <li {{if is_active .current_url "/admin/category"}}class="active"{{end}}>
                    <a href="/admin/category"><i class="fa fa-dashboard fa-fw"></i> 分类管理</a>
                </li>
                <li {{if is_active .current_url "/admin/file"}}class="active"{{end}}>
                    <a href="/admin/file"><i class="fa fa-briefcase fa-fw"></i> 资料管理</a>
                </li>
                <li {{if is_active .current_url "/admin/user"}}class="active"{{end}}>
                    <a href="/admin/user"><i class="fa fa-user fa-fw"></i> 用户管理</a>
                </li>
                <li {{if is_active .current_url "/admin/picture"}}class="active"{{end}}>
                    <a href="/admin/picture"><i class="fa fa-picture-o fa-fw"></i> 图片库</a>
                </li>
                <li {{if is_active .current_url "/admin/config"}}class="active"{{end}}>
                    <a href="/admin/config"><i class="fa fa-cog fa-fw"></i> 基础设置</a>
                </li>
            </ul>
        </div>
    </div>
    <div id="right-content-box" class="right-content-box">
        {{.LayoutContent}}
    </div>
</div>

<!-- Scripts -->
<script src="/static/js/bootstrap.min.js"></script>
<script src="/static/js/custom.js"></script>
</body>
</html>