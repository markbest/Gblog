<div class="panel panel-default search-bar">
    <form class="form-inline clearfix" method="get" id="searchform" action="/search">
        <input class="form-control" type="text" name="s" value="{{.s}}" placeholder="搜索...">
        <button type="submit" class="btn btn-danger btn-small"><i class="fa fa-search"></i></button>
    </form>
</div>
<div class="panel panel-default">
    <div class="panel-heading"><i class="fa fa-leaf"></i> 最新文章</div>
    <ul class="list-group list-group-flush">
        {{range $id, $article := .latest}}
        <li class="list-group-item">
            <a href="/article/{{$article.Id}}" title="{{$article.Title}}">
                {{$article.Title}}
            </a>
        </li>
        {{end}}
    </ul>
</div>
<div class="panel panel-default">
    <div class="panel-heading"><i class="fa fa-fire"></i> 最热文章</div>
    <ul class="list-group list-group-flush">
        {{range $id, $article := .hot}}
        <li class="list-group-item">
            <a href="/article/{{$article.Id}}" title="{{$article.Title}}">
                {{$article.Title}}
            </a>
            <label class="badge">{{$article.Views}}</label>
        </li>
        {{end}}
    </ul>
</div>
<div class="panel panel-default">
    <div class="panel-heading"><i class="fa fa-tags"></i> 热门标签</div>
    <div class="tags-list-content">
        {{range $id, $tags := .tags}}
            {{range $tag, $tid := $tags}}
            <a href="/article/{{$tid}}">{{$tag}}</a>
            {{end}}
        {{end}}
    </div>
</div>
<script>
    $(document).ready(function(){
        /*多彩tag*/
        var tags_a = $(".tags-list-content").find("a");
        tags_a.each(function(){
            var x = 9;
            var y = 0;
            var rand = parseInt(Math.random() * (x - y + 1) + y);
            $(this).addClass("size"+rand);
        });
    });
</script>