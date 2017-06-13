<div class="container">
    <div class="row">
        <div class="col-lg-12">
            <div class="row">
                <div class="col-sm-12">
                    <div class="demo-wrapper">
                        <div class="admin-page-header">
                            <i class="fa fa-home fa-fw"></i>Drag and Drop Picture <a class="pic_oper" href="/admin/picture/edit">View All Picture</a>
                        </div>
                        <div class="row demo-columns">
                            <div class="col-md-6">
                                <div id="drag-and-drop-zone" class="uploader">
                                    <div>Drag &amp; Drop Images Here</div>
                                    <div class="or">-or-</div>
                                    <div class="browser">
                                        <label>
                                            <span>Click to open the file Browser</span>
                                            <input type="hidden" name="_xsrf" value="{{.xsrf_token}}" />
                                            <input type="file" name="file" accept="image/*" multiple="multiple" title='Click to add Images'>
                                        </label>
                                    </div>
                                </div>
                                <div class="panel panel-default">
                                    <div class="panel-heading">
                                        <h3 class="panel-title">Debug</h3>
                                    </div>
                                    <div class="panel-body demo-panel-debug">
                                        <ul id="demo-debug">
                                        </ul>
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-6">
                                <div class="panel panel-default">
                                    <div class="panel-heading">
                                        <h3 class="panel-title">Uploads</h3>
                                    </div>
                                    <div class="panel-body demo-panel-files" id='demo-files'>
                                        <span class="demo-note">No Files have been selected/droped yet...</span>
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-12">
                                <div class="panel panel-default">
                                    <div class="panel-heading">
                                        <h3 class="panel-title">Uploads Review</h3>
                                    </div>
                                    <div class="panel-body" id="uploads_review"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <link rel="stylesheet" href="/static/css/uploader.css" />
                    <script type="text/javascript" src="/static/js/demo-preview.js"></script>
                    <script type="text/javascript" src="/static/js/dmuploader.js"></script>
                    <script type="text/javascript">
                        $('#drag-and-drop-zone').dmUploader({
                            url: '/admin/picture/upload',
                            dataType: 'json',
                            extraData:{'_xsrf': $('input[name="_xsrf"]').val()},
                            allowedTypes: 'image/*',
                            onInit: function(){
                                $.danidemo.addLog('#demo-debug', 'default', 'Plugin initialized correctly');
                            },
                            onBeforeUpload: function(id){
                                $.danidemo.addLog('#demo-debug', 'default', 'Starting the upload of #' + id);
                                $.danidemo.updateFileStatus(id, 'default', 'Uploading...');
                            },
                            onNewFile: function(id, file){
                                $.danidemo.addFile('#demo-files', id, file);
                                if(typeof FileReader !== "undefined"){
                                    var reader = new FileReader();
                                    var img = $('#demo-files').find('.demo-image-preview').eq(0);
                                    reader.onload = function (e){
                                        img.attr('src', e.target.result);
                                    }
                                    reader.readAsDataURL(file);
                                }else{
                                    $('#demo-files').find('.demo-image-preview').remove();
                                }
                            },
                            onComplete: function(){
                                $.danidemo.addLog('#demo-debug', 'default', 'All pending tranfers completed');
                            },
                            onUploadProgress: function(id, percent){
                                var percentStr = percent + '%';
                                $.danidemo.updateFileProgress(id, percentStr);
                            },
                            onUploadSuccess: function(id, data){
                                $.danidemo.addLog('#demo-debug', 'success', 'Upload of file #' + id + ' completed');
                                $.danidemo.addLog('#demo-debug', 'info', 'Server Response for file #' + id + ': ' + JSON.stringify(data));
                                $.danidemo.updateFileStatus(id, 'success', 'Upload Complete');
                                $.danidemo.updateFileProgress(id, '100%');

                                $('#uploads_review').append('<img src="/static/uploads/'+data['message']+'" style="width:100px;height:100px;"/>');
                            },
                            onUploadError: function(id, message){
                                $.danidemo.updateFileStatus(id, 'error', message);
                                $.danidemo.addLog('#demo-debug', 'error', 'Failed to Upload file #' + id + ': ' + message);
                            },
                            onFileTypeError: function(file){
                                $.danidemo.addLog('#demo-debug', 'error', 'File \'' + file.name + '\' cannot be added: must be an image');
                            },
                            onFileSizeError: function(file){
                                $.danidemo.addLog('#demo-debug', 'error', 'File \'' + file.name + '\' cannot be added: size excess limit');
                            },
                            onFallbackMode: function(message){
                                $.danidemo.addLog('#demo-debug', 'info', 'Browser not supported(do something else here!): ' + message);
                            }
                        });
                    </script>
                </div>
            </div>
        </div>
    </div>
</div>