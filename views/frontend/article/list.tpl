<div id="category_article_list">
    {{range $id, $article := .article}}
    <div class="list_content well">
        <i class="fa fa-bookmark fa-3x article-stick visible-md visible-lg"></i>
        <div class="data-article">
            <span class="month">{{date $article.Created_at "m"}}月</span>
            <span class="day">{{date $article.Created_at "d"}}日</span>
        </div>
        <div class="title-article">
            <h1>
                <a title="{{$article.Title}}" href="/article/{{$article.Id}}">{{$article.Title}}</a>
            </h1>
        </div>
        <div class="tag-article">
            <span class="label"><i class="fa fa-tags"></i> {{$article.Slug}}</span>
            <span class="label"><i class="fa fa-user"></i> {{$article.User.Name}}</span>
            <span class="label"><i class="fa fa-eye"></i> {{$article.Views}}</span>
        </div>
        <div class="short_content">{{$article.Summary}}</div>
        <div class="article_addition">
            <a class="btn btn-danger pull-right read-more" href="article/{{$article.Id}}" title="详细阅读 {{$article.Title}}">
                阅读全文 <span class="badge">{{$article.Views}}</span>
            </a>
        </div>
    </div>
    {{end}}
</div>
<div class="front_page">
    <div class="col-sm-12" style="padding:0px;">
        <div class="pages_content">
            {{str2html .page}}
        </div>
    </div>
</div>