package router

import (
	"gf-app/app/api/ai"
	"gf-app/boot"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func Auth(r *ghttp.Request)  {
	//Todo:统一管理消息
	if rq , err := boot.Client.Header(
		g.MapStrStr{
			"Authorization":r.Header.Get("Authorization"),
		}).Get("http://token/auth"); err != nil{
		r.Response.WriteJsonExit(g.Map{"code":"500","msg":"server error","detail":err.Error()})
	} else if rq.ReadAllString() != "ok"{
		glog.Info(rq.ReadAllString())
		r.Response.WriteJsonExit(g.Map{"code":"403","msg":"forbidden","detail":"请先登录获取token"})
	} else {
		r.Middleware.Next()
	}
}

func init() {
	s := g.Server()
	s.Use(Auth)
	s.BindHandler("/ai/*pth",ai.Ai)
}
