<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <h3 class="admin-page-header"><i class="fa fa-home fa-fw"></i>文章管理 / 新增文章</h3>
        </div>
    </div>
    <div class="row">
        <div class="col-md-12">
            <div class="admin-panel-body">
                <form action="/admin/article/create" method="POST">
                    <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
                    <div class="form-group">
                        <label>标题：</label>
                        <input type="text" name="title" class="form-control" required="required" value="">
                    </div>
                    <div class="form-group">
                        <label>关键字：</label>
                        <input type="text" name="slug" class="form-control" required="required" value="">
                    </div>
                    <div class="form-group">
                        <label>分类：</label>
                        <select name="cat_id" class="form-control" required="required">
                            <option value=""></option>
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
                        <label>短描述：</label>
                        <textarea class="form-control" rows="3" name="summary"></textarea>
                    </div>
                    <div class="form-group">
                        <label>文章内容：</label>
                        <div id="editor">
                            <textarea style="display:none;" name="body"></textarea>
                        </div>
                    </div>
                    <button type="submit" style="font-size:12px;padding:4px 10px;margin-top:-20px;" class="btn btn-primary"><i class="fa fa-floppy-o fa-fw"></i>保存</button>
                </form>
            </div>
        </div>
    </div>
</div>
<link rel="stylesheet" href="/static/css/editormd.css" />
<script src="/static/js/editormd.min.js"></script>
<script type="text/javascript">
    $(function() {
        editormd("editor", {
            width   : "100%",
            height  : 640,
            syncScrolling : "single",
            path    : "/static/js/lib/",
            imageUpload    : true,
            imageFormats   : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
            imageUploadURL : "/admin/markdown/upload?_xsrf=" + $('input[name="_xsrf"]').val(),
        });
    });
</script>