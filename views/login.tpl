<div class="container">
    <div class="row">
        <div class="col-md-4 col-md-offset-4">
            <div class="login-logo">
                <img src="/static/images/logo.png" />
            </div>
            <div class="panel panel-default">
                <div class="panel-heading">用户登录</div>
                <div class="panel-body">
                    <form role="form" method="POST" action="/customer/login">
                        <input type="hidden" name="_xsrf" value="{{.xsrf_token}}" />
                        <p>{{.flash.error}}</p>
                        <div class="form-group">
                            <input type="email"  placeholder="E-mail" class="form-control" name="email" required="required" autofocus value="">
                        </div>
                        <div class="form-group">
                            <input type="password" placeholder="Password" class="form-control" required="required" name="password">
                        </div>
                        <div class="form-group">
                            <button type="submit" class="btn btn-primary">登录</button>
                            <a href="/customer/register"><button type="button" class="btn btn-primary">注册</button></a>
                            <a class="btn btn-link" href="/">返回首页</a>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>