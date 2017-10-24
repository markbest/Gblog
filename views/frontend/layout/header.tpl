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
                            <a href="/category/{{$category.Title}}">{{$category.Title}}</a>
                        {{end}}
                    </li>
                    {{end}}
                </ul>
                <ul class="nav navbar-nav navbar-right" style="margin-right:0px;">
                    <li class="custom-button"><a href="http://static.markbest.site">静态站点</a></li>
                    <li class="custom-button"><a href="/works">开源项目</a></li>
                </ul>
            </div>
        </div>
    </nav>
</div>