package ai

import (
	"gf-app/boot"
	"gf-app/lib/resp"
	"github.com/gogf/gf/net/ghttp"
)

// Ai 动态路由解析
func Ai(r *ghttp.Request) {
	if rq, err := boot.Client.Post("http://"+r.GetString("pth"), r.GetBody()); err != nil {
		r.Response.WriteStatusExit(404,resp.Resp404(err.Error()))
	} else {
		r.Response.WriteExit(rq.ReadAll())

	}
}



