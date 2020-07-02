package ai

import (
	"gf-app/boot"
	"gf-app/lib/resp"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

func SendMessage(data string)  {
	if rq, err := boot.Client.Header(map[string]string{"Content-Type":"application/json;charset=utf-8"}).Post("https://oapi.dingtalk.com/robot/send?access_token=da896c336818972f34a626c01af1a6e4063ef6fd0ba80b25907d6e14f864f4f0",
		data); err != nil {
		glog.Warning("MessageHook错误")
	} else {
		glog.Info(rq.ReadAllString())
	}

}

// Ai 动态路由解析
func Ai(r *ghttp.Request) {
	if rq, err := boot.Client.Post("http://"+r.GetString("pth"), r.GetBody()); err != nil {
		r.Response.WriteStatusExit(404,resp.Resp404(err.Error()))
	} else {
		con := rq.ReadAll()
		//glog.Info(string(con))
		glog.Info(con)
		go SendMessage(`{
   "actionCard": {
       "title": "【鹦鹉识别】",
       "text": "![bird](/`+gtime.TimestampStr()+`.png)
### 鹦鹉识别结果` +
			string(con) + `",
   },
   "msgtype": "actionCard"
}`)
		r.Response.Write(con)
	}
}



