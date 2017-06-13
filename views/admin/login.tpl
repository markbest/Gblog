<div class="row">
    <div class="col-md-4 col-md-offset-4">
        <div class="login-logo">
            <img src="/static/images/logo.png" />
        </div>
        <div class="login-panel panel panel-default">
            <div class="panel-heading">
                <h3 class="panel-title">Please Sign In</h3>
            </div>
            <div class="panel-body">
                <form role="form" method="POST" action="/admin/login">
                    <input type="hidden" name="_xsrf" value="{{.xsrf_token}}" />
                    <fieldset>
                        <div class="form-group">
                            <input class="form-control" placeholder="E-mail" name="email" required="required" type="email" autofocus value="">
                        </div>
                        <div class="form-group">
                            <input class="form-control" placeholder="Password" name="password" required="required" type="password" value="">
                        </div>
                        <button class="btn btn-success btn-block">Login</button>
                    </fieldset>
                </form>
            </div>
        </div>
    </div>
</div>