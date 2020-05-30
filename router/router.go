package router

import (
	"gf-app/app/api/ai"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func Auth(r *ghttp.Request)  {
	//Todo:鉴权中间件
	glog.Info("权限认证")
	r.Middleware.Next()
}
func init() {
	s := g.Server()
	s.Group("/ai", func(group *ghttp.RouterGroup) {
		group.Middleware(Auth)
		group.ALL("/:name", ai.Ai)
	})
}
