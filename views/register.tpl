<div class="container">
    <div class="row">
        <div class="col-md-4 col-md-offset-4">
            <div class="login-logo">
                <img src="/static/images/logo.png" />
            </div>
            <div class="panel panel-default">
                <div class="panel-heading">用户注册</div>
                <div class="panel-body">
                    <form role="form" method="POST" action="/customer/register">
                        <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
                        <p>{{.flash.error}}</p>
                        <div class="form-group">
                            <input type="text" placeholder="昵称" class="form-control" name="name" required="required" value="">
                        </div>
                        <div class="form-group">
                            <input type="email" placeholder="邮箱" class="form-control" name="email" required="required" value="">
                        </div>

                        <div class="form-group">
                            <input type="password" placeholder="密码" class="form-control" required="required" name="password">
                        </div>

                        <div class="form-group">
                            <input type="password" placeholder="确认密码" class="form-control" required="required" name="repassword">
                        </div>
                        <div class="form-group">
                            <button type="submit" class="btn btn-primary">注册</button>
                            <a class="btn btn-link" href="/">返回首页</a>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>