package router

import (
	"gf-app/utils/ai"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Configs []ai.Config


func init() {
	s := g.Server()

	//url全小写
	s.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)

	//取出配置
	if configs , ok := g.Cfg().Get("ai").(Configs); ok {

	//AI组 post{"image":base64}(config) return {"name":"","score":""}(fetch)
	s.Group("/ai", func(group *ghttp.RouterGroup) {
		for _, config := range configs {
			client := config.Client()
			s.BindObject("/" + config.Url + "/", client)
		}
	})
	}

}
