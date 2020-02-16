package hello

import (
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/os/glog"
)

// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
    request_json, err := r.GetJson()
    if err != nil {
        glog.Error(err)
        return
    }


    r.Response.Writeln("Hello World!")
}
