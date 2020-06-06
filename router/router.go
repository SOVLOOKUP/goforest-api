package router

import (
	"gf-app/app/api/ai"
	"gf-app/boot"
	"gf-app/lib/resp"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)


func Auth(r *ghttp.Request)  {


	if rq , err := boot.Client.Header(
		g.MapStrStr{
			"Authorization":r.Header.Get("Authorization"),
		}).Get("http://token/auth"); err != nil{
		r.Response.WriteStatusExit(500,resp.Resp500(err.Error()))
	} else if rq.ReadAllString() != "ok"{
		r.Response.WriteStatusExit(403,resp.Resp403("请先登录获取token"))
	} else {
		r.Middleware.Next()
	}
}

func init() {
	s := g.Server()
	s.Use(Auth)
	s.BindHandler("/ai/*pth",ai.Ai)
}
