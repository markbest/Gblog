<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <div class="admin-page-header"><i class="fa fa-home fa-fw"></i>用户管理</div>
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
                            <th class="center col-sm-2">姓名</th>
                            <th class="center col-sm-2">邮箱</th>
                            <th class="center col-sm-1">创建时间</th>
                            <th class="center col-sm-1">操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $id, $user := .users}}
                        <tr>
                            <td>{{$user.Id}}</td>
                            <td class="center">{{$user.Name}}</td>
                            <td class="center">{{$user.Email}}</td>
                            <td class="center">{{date $user.Created_at "Y-m-d H:i:s"}}</td>
                            <td class="center">
                                <a href="/admin/user/{{$user.Id}}" class="btn btn-success"><i class="fa fa-pencil fa-fw"></i>编辑</a>
                                <form action="/admin/user/{{$user.Id}}" method="POST" style="display:inline;">
                                    <input name="_method" type="hidden" value="DELETE">
                                    <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                    <button type="button" class="btn btn-danger del_btn"><i class="fa fa-trash fa-fw"></i>删除</button>
                                </form>
                            </td>
                        </tr>
                        {{end}}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
    <div class="page_html" style="text-align:right;">
        {{str2html .page}}
    </div>
</div>