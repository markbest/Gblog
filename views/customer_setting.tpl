<div class="col-lg-12">
    <div class="row">
        <div class="aw-content-wrap">
            <div class="aw-user-setting">
                <div class="tabbable">
                    <ul class="nav nav-tabs aw-nav-tabs active">
                        <li><a href="javascript::void(0);">安全设置</a></li>
                        <li class="active"><a href="javascript::void(0);">基本资料</a></li>
                        <h2>用户设置</h2>
                    </ul>
                </div>
                <div class="tab-content" style="display:none;">
                    <div class="aw-mod">
                        <div class="mod-body">
                            <div class="aw-mod mod-base">
                                <div class="mod-head">
                                    <h3>安全信息</h3>
                                </div>
                                <div class="mod-body" style="height:100%;overflow:hidden;">
                                    <form id="setting_form" method="post" action="/customer/setting">
                                        <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
                                        <input type="hidden" name="customer_id" value="{{.loginCustomer.Id}}">
                                        <p>{{.flash.error}}</p>
                                        <dl><dt>初始密码:</dt><dd><input class="form-control" type="password" name="old_password" required="required"></dd></dl>
                                        <dl><dt>新改密码:</dt><dd><input class="form-control" type="password" name="password" required="required"></dd></dl>
                                        <dl><dt>确认密码:</dt><dd><input class="form-control" type="password" name="password_confirmation" required="required"></dd></dl>
                                        <input type="submit" class="setting_customer_sub" value="保存">
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="tab-content">
                    <div class="aw-mod">
                        <div class="mod-body">
                            <div class="aw-mod mod-base">
                                <div class="mod-head">
                                    <h3>基本信息</h3>
                                </div>
                                <div class="mod-body" style="height:100%;overflow:hidden;">
                                    <form id="setting_form" method="post" action="/customer/setting">
                                        <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
                                        <input type="hidden" name="customer_id" value="15">
                                        <dl><dt>账号:</dt><dd>{{.loginCustomer.Email}}</dd></dl>
                                        <dl><dt>真实姓名:</dt><dd><input class="form-control" type="text" value="{{.loginCustomer.Name}}" name="name" required="required"></dd></dl>
                                        <input type="submit" class="setting_customer_sub" value="保存">
                                    </form>
                                    <div class="side-bar">
                                        <dl>
                                            <dt class="pull-left">
                                                <img id="return_customer_icon" src="/static/uploads/{{.loginCustomer.Icon}}" width="100px" height="100px">
                                            </dt>
                                            <dd class="pull-left">
                                                <h5>头像设置(jpg、gif、png)</h5>
                                                <p>
                                                    <input style="line-height:20px;width:220px;" id="customer_icon" type="file" size="20" name="customer_icon" class="input">
                                                    <button id="buttonUpload">上传图像</button>
                                                    <span id="avatar_uploading_status" class="hide"><i class="aw-loading"></i> 文件上传中...</span>
                                                </p>
                                            </dd>
                                        </dl>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<script type="text/javascript" src="/static/js/ajaxfileupload.js"></script>
<script language="javascript">
    jQuery(function(){
        $("#buttonUpload").click(function(){
            $('#avatar_uploading_status').removeClass('hide');
            $.ajaxFileUpload({
                url:'/customer/upload',
                secureuri :false,
                fileElementId :'customer_icon',
                data:{
                    '_xsrf': $('input[name="_xsrf"]').val(),
                    'id': $('input[name="customer_id"').val()
                },
                success : function (data, status){
                    $('#avatar_uploading_status').addClass('hide');
                    $('.customer_small_icon').attr('src', data.message);
                    $('#return_customer_icon').attr('src', data.message).show();
                    location.reload();
                },
                error: function(data, status, e){
                    console.log(e)
                }
            });
        });

        $('.aw-nav-tabs li').each(function(){
            $(this).click(function(){
                var index = $(this).index();
                $('.aw-nav-tabs li').removeClass('active');
                $(this).addClass('active');
                $('.tab-content').hide();
                $('.tab-content').eq(index).show();
            });
        });
    });
</script>