package ai

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//client 连接池
var client = g.Client()

// Ai 动态路由解析
func Ai(r *ghttp.Request) {

	if rq, err := client.Post("http://"+r.GetString("pth"), r.GetBody()); err != nil {
		r.Response.WriteJsonExit(g.Map{"code":"0","msg":"router found","detail":err.Error()})
	} else {
		r.Response.WriteJsonExit(rq.ReadAll())
	}

}



