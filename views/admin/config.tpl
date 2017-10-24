<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <div class="admin-page-header">
                <div class="col-sm-4"><i class="fa fa-home fa-fw"></i>基础设置</div>
                <div class="col-sm-8 align-right">
                    <button class="admin-btn-head btn btn-primary" data-toggle="modal" data-target="#add_setting">
                        <i class="fa fa-plus-circle fa-fw"></i>新增设置
                    </button>
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
                            <th class="col-sm-2">配置标题</th>
                            <th class="col-sm-1">配置路径</th>
                            <th class="col-sm-4">配置值</th>
                            <th class="col-sm-1">创建时间</th>
                            <th class="col-sm-2"></th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $id, $config := .configs}}
                        <tr>
                            <td>{{$config.Id}}</td>
                            <td>{{$config.Name}}</td>
                            <td>{{$config.Path}}</td>
                            <td>{{$config.Value}}</td>
                            <td>{{date $config.Created_at "m-d H:i"}}</td>
                            <td class="center">
                                <button type="button" data-toggle="modal" data-target="#edit_config_{{$config.Id}}" class="admin-btn btn btn-success">
                                    <i class="fa fa-pencil fa-fw"></i>编辑</a>
                                </button>
                                <form action="/admin/config/{{$config.Id}}" method="POST" style="display: inline;">
                                    <input name="_method" type="hidden" value="DELETE">
                                    <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                    <button type="button" class="admin-btn btn btn-danger del_btn"><i class="fa fa-trash fa-fw"></i>删除</button>
                                </form>
                            </td>
                            <div class="modal fade" id="edit_config_{{$config.Id}}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="display: none;">
                                <div class="modal-dialog">
                                    <div class="modal-content">
                                        <div class="modal-header">
                                            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                                            <h4 class="modal-title" id="myModalLabel">{{$config.Name}}</h4>
                                        </div>
                                        <form action="/admin/config/{{$config.Id}}" method="post">
                                            <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                            <div class="modal-body">
                                                <div class="form-group">
                                                    <label>配置标题：</label>
                                                    <input type="text" name="name" class="form-control" required="required" value="{{$config.Name}}">
                                                </div>
                                                <div class="form-group">
                                                    <label>配置路径：</label>
                                                    <input type="text" name="path" class="form-control" required="required" value="{{$config.Path}}">
                                                </div>
                                                <div class="form-group">
                                                    <label>配置值：</label>
                                                    <input type="text" name="value" class="form-control" required="required" value="{{$config.Value}}">
                                                </div>
                                            </div>
                                            <div class="modal-footer">
                                                <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                                                <button type="submit" class="btn btn-primary">保存</button>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                            </div>
                        </tr>
                        {{end}}
                        </tbody>
                    </table>
                    <div class="modal fade" id="add_setting" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="display: none;">
                        <div class="modal-dialog">
                            <div class="modal-content">
                                <div class="modal-header">
                                    <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                                    <h4 class="modal-title" id="myModalLabel">新增设置</h4>
                                </div>
                                <form action="/admin/config" method="post">
                                    <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
                                    <div class="modal-body">
                                        <div class="form-group">
                                            <label>配置标题：</label>
                                            <input type="text" name="name" class="form-control" required="required" value="">
                                        </div>
                                        <div class="form-group">
                                            <label>配置路径：</label>
                                            <input type="text" name="path" class="form-control" required="required" value="">
                                        </div>
                                        <div class="form-group">
                                            <label>配置值：</label>
                                            <input type="text" name="value" class="form-control" required="required" value="">
                                        </div>
                                    </div>
                                    <div class="modal-footer">
                                        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                                        <button type="submit" class="btn btn-primary">保存</button>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>