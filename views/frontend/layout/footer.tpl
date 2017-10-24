<div id="footer">
    {{index .configs "web_copyright"}}
</div>
<a id="scrollUp"></a>
<script type="text/javascript">
    jQuery(function($){
        $(window).scroll(function(){
            var scrollt = document.documentElement.scrollTop + document.body.scrollTop;
            if(scrollt >200){
                $("#scrollUp").fadeIn(400);
            }else{
                $("#scrollUp").stop().fadeOut(400);
            }
        });
        $("#scrollUp").click(function(){
            $("html,body").animate({scrollTop:"0px"},200);
        });
    });
</script>