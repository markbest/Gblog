<div class="top_content">
    <nav class="navbar navbar-default container">
        <div class="container-fluid">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#as-example-navbar-collapse-1">
                    <span class="sr-only">Toggle Navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <a class="navbar-brand logo" title="你的指尖有改变世界的力量 - markbest.site" href="/">
                    <img src="/static/images/logo.png" />
                </a>
            </div>
            <div class="collapse navbar-collapse" id="as-example-navbar-collapse-1">
                <ul class="nav navbar-nav">
                    <li {{if eq .current_url "/"}}class="active"{{end}}>
                        <a href="/">首页</a>
                    </li>

                    {{range $id, $category := .category}}
                    <li {{if is_active $.current_url $category.Title}}class="active"{{end}}>
                        {{if $category.Sub_category}}
                        <a data-toggle="dropdown" href="#">{{$category.Title}} <span class="caret"></span></a>
                        <ul class="dropdown-menu">
                            {{range $id, $category := $category.Sub_category}}
                            <li><a href="/category/{{$category.Title}}">{{$category.Title}}</a></li>
                            {{end}}
                        </ul>
                        {{else}}
                            {{if eq $category.Title "资料下载"}}
                            <a href="/category/files">{{$category.Title}}</a>
                            {{else}}
                            <a href="/category/{{$category.Title}}">{{$category.Title}}</a>
                            {{end}}
                        {{end}}
                    </li>
                    {{end}}
                </ul>
                <ul class="nav navbar-nav navbar-right" style="margin-right:0px;">
                    <li class="customer-login"><a href="/works">作品</a></li>
                    {{if .isLogin}}
                    <li class="dropdown">
                        <a href="###" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
                            <img class="customer_small_icon" width="25px" height="25px" src="/static/uploads/{{.loginCustomer.Icon}}">
                            {{.loginCustomer.Name}} <span class="caret"></span>
                        </a>
                        <ul class="dropdown-menu" role="menu">
                            <li><a href="/customer/home">我的信息</a></li>
                            <li><a href="/customer/logout">退出登录</a></li>
                        </ul>
                    </li>
                    {{else}}
                    <li class="customer-login"><a href="/customer/login">登录</a></li>
                    <li class="customer-register"><a href="/customer/register">注册</a></li>
                    {{end}}
                </ul>
            </div>
        </div>
    </nav>
</div>