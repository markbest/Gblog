<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <div class="admin-page-header">
                <div class="col-sm-4">
                    <i class="fa fa-home fa-fw"></i>分类管理
                </div>
                <div class="col-sm-8 align-right">
                    <button type="button" data-toggle="modal" data-target="#add_category" class="admin-btn-head btn btn-primary">
                        <i class="fa fa-plus-circle fa-fw"></i>新增分类
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
                            <th class="col-sm-4">标题</th>
                            <th class="center col-sm-1">文章数</th>
                            <th class="center col-sm-2">创建时间</th>
                            <th class="center col-sm-1">排序</th>
                            <th class="center col-sm-2">操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $id, $category := .category}}
                        <tr>
                            <td class="level-1">{{$category.Title}}</td>
                            <td class="center">{{$category.Count_articles}}</td>
                            <td class="center"><span class="updated_at">{{date $category.Created_at "Y-m-d H:i:s"}}</span></td>
                            <td class="center">{{$category.Sort}}</td>
                            <td class="center">
                                <button type="button" data-toggle="modal" data-target="#edit_category_{{$category.Id}}" class="admin-btn btn btn-success">
                                    <i class="fa fa-pencil fa-fw"></i>编辑</a>
                                </button>
                                <form action="/admin/category/{{$category.Id}}" method="POST" style="display:inline;">
                                    <input name="_method" type="hidden" value="DELETE">
                                    <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                    <button type="button" class="btn btn-danger del_btn"><i class="fa fa-trash fa-fw"></i>删除</button>
                                </form>
                            </td>
                        </tr>
                        <div class="modal fade" id="edit_category_{{$category.Id}}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="display:none;">
                            <div class="modal-dialog">
                                <div class="modal-content">
                                    <div class="modal-header">
                                        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                                        <h4 class="modal-title" id="myModalLabel">编辑分类</h4>
                                    </div>
                                    <form action="/admin/category/{{$category.Id}}" method="POST">
                                        <input name="_method" type="hidden" value="PUT">
                                        <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                        <div class="modal-body">
                                            <div class="form-group">
                                                <label>分类名称：</label>
                                                <input type="text" name="title" class="form-control" required="required" value="{{$category.Title}}">
                                            </div>
                                            <div class="form-group">
                                                <label>分类归属：</label>
                                                <select name="parent_id" class="form-control" required="required">
                                                    <option value="0">顶级分类</option>
                                                    {{range $id, $s_category := $.category}}
                                                    <option value="{{$s_category.Id}}">{{$s_category.Title}}</option>
                                                    {{if $s_category.Sub_category}}
                                                    {{range $id, $ss_category := $s_category.Sub_category}}
                                                    <option value="{{$ss_category.Id}}">&nbsp;&nbsp;&nbsp;&nbsp;{{$ss_category.Title}}</option>
                                                    {{end}}
                                                    {{end}}
                                                    {{end}}
                                                </select>
                                            </div>
                                            <div class="form-group">
                                                <label>分类排序：</label>
                                                <input type="text" name="sort" class="form-control" required="required" value="{{$category.Sort}}">
                                            </div>
                                        </div>
                                        <div class="modal-footer">
                                            <button type="button" class="admin-btn btn btn-default" data-dismiss="modal"><i class="fa fa-ban"></i>关闭</button>
                                            <button type="submit" class="admin-btn btn btn-primary"><i class="fa fa-floppy-o"></i>保存</button>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>
                        {{if $category.Sub_category}}
                            {{range $id, $category := $category.Sub_category}}
                            <tr>
                                <td class="level-2">{{$category.Title}}</td>
                                <td class="center">{{$category.Count_articles}}</td>
                                <td class="center"><span class="updated_at">{{date $category.Created_at "Y-m-d H:i:s"}}</span></td>
                                <td class="center">{{$category.Sort}}</td>
                                <td class="center">
                                    <button type="button" data-toggle="modal" data-target="#edit_category_{{$category.Id}}" class="admin-btn btn btn-success">
                                        <i class="fa fa-pencil fa-fw"></i>编辑</a>
                                    </button>
                                    <form action="/admin/category/{{$category.Id}}" method="POST" style="display:inline;">
                                        <input name="_method" type="hidden" value="DELETE">
                                        <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                        <button type="button" class="btn btn-danger del_btn"><i class="fa fa-trash fa-fw"></i>删除</button>
                                    </form>
                                </td>
                            </tr>
                            <div class="modal fade" id="edit_category_{{$category.Id}}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="display:none;">
                                <div class="modal-dialog">
                                    <div class="modal-content">
                                        <div class="modal-header">
                                            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                                            <h4 class="modal-title" id="myModalLabel">编辑分类</h4>
                                        </div>
                                        <form action="/admin/category/{{$category.Id}}" method="POST">
                                            <input name="_method" type="hidden" value="PUT">
                                            <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                            <div class="modal-body">
                                                <div class="form-group">
                                                    <label>分类名称：</label>
                                                    <input type="text" name="title" class="form-control" required="required" value="{{$category.Title}}">
                                                </div>
                                                <div class="form-group">
                                                    <label>分类归属：</label>
                                                    <select name="parent_id" class="form-control" required="required">
                                                        <option value="0">顶级分类</option>
                                                        {{range $id, $s_category := $.category}}
                                                        <option {{if eq $category.Parent_id $s_category.Id}} selected="selected" {{end}} value="{{$s_category.Id}}">{{$s_category.Title}}</option>
                                                        {{if $s_category.Sub_category}}
                                                        {{range $id, $ss_category := $s_category.Sub_category}}
                                                        <option value="{{$ss_category.Id}}">&nbsp;&nbsp;&nbsp;&nbsp;{{$ss_category.Title}}</option>
                                                        {{end}}
                                                        {{end}}
                                                        {{end}}
                                                    </select>
                                                </div>
                                                <div class="form-group">
                                                    <label>分类排序：</label>
                                                    <input type="text" name="sort" class="form-control" required="required" value="{{$category.Sort}}">
                                                </div>
                                            </div>
                                            <div class="modal-footer">
                                                <button type="button" class="admin-btn btn btn-default" data-dismiss="modal"><i class="fa fa-ban"></i>关闭</button>
                                                <button type="submit" class="admin-btn btn btn-primary"><i class="fa fa-floppy-o"></i>保存</button>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                            </div>
                            {{end}}
                        {{end}}
                        {{end}}
                        </tbody>
                    </table>
                    <div class="modal fade" id="add_category" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="display: none;">
                        <div class="modal-dialog">
                            <div class="modal-content">
                                <div class="modal-header">
                                    <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                                    <h4 class="modal-title" id="myModalLabel">新增分类</h4>
                                </div>
                                <form action="/admin/category" method="POST">
                                    <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
                                    <div class="modal-body">
                                        <div class="form-group">
                                            <label>分类名称：</label>
                                            <input type="text" name="title" class="form-control" required="required" value="">
                                        </div>
                                        <div class="form-group">
                                            <label>分类归属：</label>
                                            <select name="parent_id" class="form-control" required="required">
                                                <option value="0">顶级分类</option>
                                                {{range $id, $category := .category}}
                                                <option value="{{$category.Id}}">{{$category.Title}}</option>
                                                {{if $category.Sub_category}}
                                                    {{range $id, $category := $category.Sub_category}}
                                                    <option value="{{$category.Id}}">&nbsp;&nbsp;&nbsp;&nbsp;{{$category.Title}}</option>
                                                    {{end}}
                                                {{end}}
                                                {{end}}
                                            </select>
                                            <input type="hidden" name="sort" value="0">
                                        </div>
                                    </div>
                                    <div class="modal-footer">
                                        <button type="button" class="admin-btn btn btn-default" data-dismiss="modal"><i class="fa fa-ban"></i>关闭</button>
                                        <button type="submit" class="admin-btn btn btn-primary"><i class="fa fa-floppy-o"></i>保存</button>
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