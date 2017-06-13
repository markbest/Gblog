<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <div class="admin-page-header">
                <div class="col-sm-4"><i class="fa fa-home fa-fw"></i>资料管理</div>
                <div class="col-sm-8 align-right">
                    <button class="admin-btn-head btn btn-primary" data-toggle="modal" data-target="#add_file">
                        <i class="fa fa-plus-circle fa-fw"></i>上传资料
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
                            <th class="col-sm-2">标题</th>
                            <th class="col-sm-2">名称</th>
                            <th class="col-sm-1">大小</th>
                            <th class="col-sm-1">分类</th>
                            <th class="col-sm-1">类型</th>
                            <th class="col-sm-1">创建时间</th>
                            <th class="col-sm-2"></th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $id, $file := .files}}
                        <tr>
                            <td>{{$file.Id}}</td>
                            <td>{{$file.Title}}</td>
                            <td>{{$file.Name}}</td>
                            <td>{{$file.Size}}</td>
                            <td>{{$file.Cat.Title }}</td>
                            <td>{{$file.Type}}</td>
                            <td>{{date $file.Created_at "m-d H:i"}}</td>
                            <td class="center">
                                <button type="button" data-toggle="modal" data-target="#edit_file_{{$file.Id}}" class="admin-btn btn btn-success">
                                    <i class="fa fa-pencil fa-fw"></i>编辑</a>
                                </button>
                                <form action="/admin/file/{{$file.Id}}" method="POST" style="display: inline;">
                                    <input name="_method" type="hidden" value="DELETE">
                                    <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                    <button type="button" class="admin-btn btn btn-danger del_btn"><i class="fa fa-trash fa-fw"></i>删除</button>
                                </form>
                            </td>
                        </tr>
                        <div class="modal fade" id="edit_file_{{$file.Id}}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="display: none;">
                            <div class="modal-dialog">
                                <div class="modal-content">
                                    <div class="modal-header">
                                        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                                        <h4 class="modal-title" id="myModalLabel">编辑资料</h4>
                                    </div>
                                    <form action="/admin/file/{{$file.Id}}" method="POST">
                                        <input name="_method" type="hidden" value="PUT">
                                        <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                        <div class="modal-body">
                                            <div class="form-group">
                                                <label>资料名称：</label>
                                                <input type="text" name="title" class="form-control" required="required" value="{{$file.Title}}">
                                            </div>
                                            <div class="form-group">
                                                <label>分类归属：</label>
                                                <select name="cat_id" class="form-control" required="required">
                                                    {{range $id, $s_category := $.category}}
                                                    <option {{if eq $file.Cat.Id $s_category.Id}} selected="selected" {{end}} value="{{$s_category.Id}}">{{$s_category.Title}}</option>
                                                    {{if $s_category.Sub_category}}
                                                    {{range $id, $ss_category := $s_category.Sub_category}}
                                                    <option value="{{$ss_category.Id}}">&nbsp;&nbsp;&nbsp;&nbsp;{{$ss_category.Title}}</option>
                                                    {{end}}
                                                    {{end}}
                                                    {{end}}
                                                </select>
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
                        </tbody>
                    </table>
                    <div class="modal fade" id="add_file" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="display:none;">
                        <div class="modal-dialog">
                            <div class="modal-content">
                                <div class="modal-header">
                                    <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                                    <h4 class="modal-title" id="myModalLabel">上传文件</h4>
                                </div>
                                <div class="col-md-12">
                                    <div id="drag-and-drop-zone" class="uploader">
                                        <div>Drag &amp; Drop Files Here</div>
                                        <div class="or">-or-</div>
                                        <div class="browser">
                                            <label>
                                                <span>Click to open the file Browser</span>
                                                <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                                                <input type="file" name="file" accept="*" multiple="multiple" title='Click to add files'>
                                            </label>
                                        </div>
                                    </div>
                                    <div class="panel panel-default" style="display:none;">
                                        <div class="panel-heading"><h3 class="panel-title">Debug</h3></div>
                                        <div class="panel-body panel-debug">
                                            <ul id="debug-container"></ul>
                                        </div>
                                    </div>
                                    <div class="panel panel-default">
                                        <div class="panel-heading"><h3 class="panel-title">Uploads</h3></div>
                                        <div class="panel-body panel-files" id='files-container'>
                                            <span class="note-container">No Files have been selected/droped yet...</span>
                                        </div>
                                    </div>
                                </div>
                                <div class="modal-footer"></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="page_html" style="text-align:right;">
        {{str2html .page}}
    </div>
</div>
<link rel="stylesheet" href="/static/css/uploader.css" />
<script src="/static/js/demo-preview.js"></script>
<script src="/static/js/dmuploader.js"></script>
<script>
    $(function(){
        $('#drag-and-drop-zone').dmUploader({
            url: '/admin/file/upload',
            dataType: 'json',
            extraData:{'_xsrf': $('input[name="_xsrf"]').val()},
            allowedTypes: '*',
            onInit: function(){
                $.danidemo.addLog('#debug-container', 'default', 'Plugin initialized correctly');
            },
            onBeforeUpload: function(id){
                $.danidemo.addLog('#debug-container', 'default', 'Starting the upload of #' + id);
                $.danidemo.updateFileStatus(id, 'default', 'Uploading...');
            },
            onNewFile: function(id, file){
                $.danidemo.addFile('#files-container', id, file);
                $('#uploader-files').find('.uploader-image-preview').remove();
            },
            onComplete: function(){
                $.danidemo.addLog('#debug-container', 'default', 'All pending tranfers completed');
                location.reload();
            },
            onUploadProgress: function(id, percent){
                var percentStr = percent + '%';
                $.danidemo.updateFileProgress(id, percentStr);
            },
            onUploadSuccess: function(id, data){
                $.danidemo.addLog('#debug-container', 'success', 'Upload of file #' + id + ' completed');
                $.danidemo.addLog('#debug-container', 'info', 'Server Response for file #' + id + ': ' + JSON.stringify(data));
                $.danidemo.updateFileStatus(id, 'success', 'Upload Complete');
                $.danidemo.updateFileProgress(id, '100%');
            },
            onUploadError: function(id, message){
                $.danidemo.updateFileStatus(id, 'error', message);
                $.danidemo.addLog('#debug-container', 'error', 'Failed to Upload file #' + id + ': ' + message);
            },
            onFileTypeError: function(file){
                $.danidemo.addLog('#debug-container', 'error', 'File \'' + file.name + '\' cannot be added: must be an image');
            },
            onFileSizeError: function(file){
                $.danidemo.addLog('#debug-container', 'error', 'File \'' + file.name + '\' cannot be added: size excess limit');
            },
            onFallbackMode: function(message){
                $.danidemo.addLog('#debug-container', 'info', 'Browser not supported(do something else here!): ' + message);
            }
        });
    });
</script>