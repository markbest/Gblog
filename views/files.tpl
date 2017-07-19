<div class="file-list-container">
    <ul>
        {{range $id, $file := .files}}
        <li>
            <h3>{{$file.Title}}</h3>
            <div class="file-info">
                <a href="/files/download/{{$file.Id}}">{{$file.Title}}</a><br>
                <span>{{date $file.Created_at "Y-m-d H:i:s"}}, {{fsize $file.Size}}</span>
            </div>
        </li>
        {{end}}
    </ul>
</div>
<div class="front_page">
    <div class="col-sm-12" style="padding:0px;">
        <div class="pages_content">
            {{str2html .page}}
        </div>
    </div>
</div>