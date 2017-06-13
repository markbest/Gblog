<div id="footer">
    Copyright © 2015 - 2017 markbest.site - 你的指尖有改变世界的力量 - All Rights Reserved.
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