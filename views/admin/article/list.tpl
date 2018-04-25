<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <div class="admin-page-header">
                <div class="col-sm-4"><i class="fa fa-home fa-fw"></i>文章管理</div>
                <div class="col-sm-8 align-right">
                    <a href="/admin/article/create" class="admin-btn-head btn btn-primary"><i class="fa fa-plus-circle fa-fw"></i>新增文章</a>
                    <form action="/admin/article" method="GET" style="display:none;">
                        <div class="input-group custom-search-form">
                            <input type="text" class="form-control" name="keywords" placeholder="Search..." value="" />
                            <span class="input-group-btn">
							    <button class="btn btn-default" type="submit" style="margin-top:-6px;">
                                    <i class="fa fa-search"></i>
                                </button>
						    </span>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
    <div class="row">
        <div class="col-lg-12">
            <div class="row">
                <div class="col-sm-12">
                    <table class="table table-striped">
                        <thead>
                        <tr>
                            <th class="col-sm-1">#</th>
                            <th class="col-sm-4">标题</th>
                            <th class="col-sm-1">标签</th>
                            <th class="center col-sm-1">分类</th>
                            <th class="center col-sm-1">创建时间</th>
                            <th class="center col-sm-1">浏览量</th>
                            <th class="center col-sm-2">操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $id, $article := .articles}}
                        <tr>
                            <td>{{$article.Id}}</td>
                            <td><a class="title" href="/admin/article/{{$article.Id}}">{{$article.Title}}</a></td>
                            <td>{{$article.Slug}}</td>
                            <td class="center"><a class="category" href="/category/{{$article.Cat.Title}}" target="_blank">{{$article.Cat.Title}}</a></td>
                            <td class="center"><span class="updated_at">{{date $article.CreatedAt "m-d H:i"}}</span></td>
                            <td class="center"><span class="updated_at">{{$article.Views}}</span></td>
                            <td class="center">
                                <a href="/admin/article/{{$article.Id}}" class="admin-btn btn btn-success"><i class="fa fa-pencil fa-fw"></i>编辑</a>
                                <form action="/admin/article/{{$article.Id}}" method="POST" style="display: inline;">
                                    <input name="_method" type="hidden" value="DELETE">
                                    <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                    <button type="button" class="admin-btn btn btn-danger del_btn"><i class="fa fa-trash fa-fw"></i>删除</button>
                                </form>
                            </td>
                        </tr>
                        {{end}}
                        </tbody>
                    </table>
                </div>
            </div>
            <div class="page_html" style="text-align:right;">
                {{str2html .page}}
            </div>
        </div>
    </div>
</div>