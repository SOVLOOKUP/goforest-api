package ai

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Ai 动态路由解析
func Ai(r *ghttp.Request) {
	// if name := genv.Get(r.GetString("name"));name != "" {
	if rq, err := ghttp.Post("http://"+r.GetString("pth"), r.GetBody()); err != nil {
		//glog.Debug(err.Error())
		r.Response.WriteJsonExit(g.Map{"code":"0","msg":"router found","detail":err.Error()})
	} else {
		defer rq.Close()
		r.Response.WriteJsonExit(rq.ReadAll())
	}
	// } else {
	// 	r.Response.WriteJsonExit(g.Map{"code":"0","msg":"router found","detail":"你来到了没有数据的地方~"})
	// }
}


