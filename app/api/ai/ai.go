package ai

import (
	"gf-app/boot"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Ai 动态路由解析
func Ai(r *ghttp.Request) {
	if rq, err := boot.Client.Post("http://"+r.GetString("pth"), r.GetBody()); err != nil {
		r.Response.WriteJsonExit(g.Map{"code":"404","msg":"router found","detail":err.Error()})
	} else {
		r.Response.WriteJsonExit(rq.ReadAll())
	}

}



