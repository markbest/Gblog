<div class="article-content">
    <div class="view">
        <ul class="breadcrumb">
            <li><i class="fa fa-home fa-fw"></i> <a href="/">主页</a>
            <li class="active">{{.article.Title}}</li>
        </ul>
    </div>
    <div class="title-article">
        <h1>
            <a title="{{.article.Title}}" href="/article/{{.article.Id}}">{{.article.Title}}</a>
        </h1>
    </div>
    <div class="tag-article">
        <span class="label"><i class="fa fa-tags"></i> {{date .article.Created_at "m-d"}}</span>
        <span class="label"><i class="fa fa-user"></i> {{.article.User.Name}}</span>
        <span class="label"><i class="fa fa-eye"></i> {{.article.Views}}</span>
    </div>
    <div id="content">
       <textarea style="display:none;">{{.article.Body}}</textarea>
    </div>
</div>
<link rel="stylesheet" href="/static/css/editormd.css" />
<script src="/static/js/lib/marked.min.js"></script>
<script src="/static/js/lib/prettify.min.js"></script>
<script src="/static/js/lib/raphael.min.js"></script>
<script src="/static/js/lib/underscore.min.js"></script>
<script src="/static/js/lib/sequence-diagram.min.js"></script>
<script src="/static/js/lib/flowchart.min.js"></script>
<script src="/static/js/lib/jquery.flowchart.min.js"></script>
<script src="/static/js/editormd.min.js"></script>
<script type="text/javascript">
    $(function() {
        editormd.markdownToHTML("content", {
            htmlDecode      : "style,script,iframe",  // you can filter tags decode
            emoji           : true,
            taskList        : true,
            tex             : true,  // 默认不解析
            flowChart       : true,  // 默认不解析
            sequenceDiagram : true,  // 默认不解析
        });
    })
</script>