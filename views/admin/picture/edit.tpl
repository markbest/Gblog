<link rel="stylesheet" href="/static/css/fancybox.css" rel="stylesheet" />
<script type="text/javascript" src="/static/js/jquery.fancybox-1.3.1.pack.js"></script>
<script type="text/javascript" src="/static/js/jquery.masonry.min.js"></script>
<script type="text/javascript">
    $(function(){
        $("a[rel=group]").fancybox({
            'transitionIn'	: 'elastic',
            'transitionOut'	: 'elastic',
            'titlePosition' : 'inside'
        });
    });
    $(document).ready(function(){
        var $container = $('.picture_wall');
        $container.imagesLoaded(function(){
            $container.masonry({
                itemSelector: '.picture_wall_list',
                columnWidth: 5 //每两列之间的间隙为5像素
            });
        });
    });
</script>
<div class="row">
    <div class="col-lg-12">
        <div class="row">
            <div class="col-sm-12">
                <ul class="picture_wall">
                    {{range $id, $picture := .pictures}}
                    <li class="picture_wall_list">
                        <form name="pic_info" action="/admin/picture/{{$picture.Id}}" method="post">
                            <input type="hidden" name="_xsrf" value="{{$.xsrf_token}}">
                            <a rel="group" href="{{base_url "/static/uploads/" $picture.Img_url}}" title="{{$picture.Note}}">
                                <img src="/static/uploads/{{$picture.Img_url}}" />
                            </a>
                            Input Picture Note :<textarea name="note" style="height:54px;width:200px;border-color:#ffffff;padding:5px;font-size:12px;">{{$picture.Note}}</textarea>
                            <input style="margin:8px 5px 0px 0px;" type="checkbox" value="1" name="is_delete" />delete <button style="padding:5px 8px;float:right;" class="admin-btn btn btn-info">Save</button>
                        </form>
                    </li>
                    {{end}}
                </ul>
            </div>
        </div>
    </div>
</div>




