package main

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/os/gview"
    "gitee.com/johng/gf/g/net/ghttp"
    "gitee.com/johng/gf/g/util/gpage"
)

func main() {
    s := ghttp.GetServer()
    s.BindHandler("/page/ajax", func(r *ghttp.Request){
        page := gpage.New(100, 10, r.Get("page"), r.URL.String(), r.Router.Uri)
        page.EnableAjax("DoAjax")
        buffer, _ := gview.ParseContent(`
        <html>
            <head>
                <style>
                    a,span {padding:8px; font-size:16px;}
                    div{margin:5px 5px 20px 5px}
                </style>
                <script src="https://cdn.bootcss.com/jquery/2.0.3/jquery.min.js"></script>
                <script>
                function DoAjax(url) {
                     $.get(url, function(data,status) {
                         $("body").html(data);
                     });
                }
                </script>
            </head>
            <body>
                <div>{{.page}}</div>
            </body>
        </html>
        `, g.Map{
            "page" : gview.HTML(page.GetContent(1)),
        })
        r.Response.Write(buffer)
    })
    s.SetPort(8199)
    s.Run()
}