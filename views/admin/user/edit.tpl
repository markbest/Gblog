<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <h3 class="admin-page-header"><i class="fa fa-home fa-fw"></i>用户管理 / {{.user.Name}}</h3>
        </div>
    </div>
    <div class="row">
        <div class="col-md-8">
            <div class="admin-panel-body">
                <form action="/admin/user/{{.user.Id}}" method="POST">
                    <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
                    <div class="form-group">
                        <label>姓名：</label>
                        <input type="text" name="name" class="form-control" required="required" value="{{.user.Name}}">
                    </div>
                    <div class="form-group">
                        <label>新密码：</label>
                        <input type="password" name="password" class="form-control" required="required" value="">
                    </div>
                    <button class="admin-btn btn btn-success">提交用户</button>
                </form>
            </div>
        </div>
    </div>
</div>