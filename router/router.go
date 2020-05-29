package router

import (
	"gf-app/utils/ai"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)
//todo:配置自动读取（分布式配置管理cli） api根据ainame动态遍历配置生成
var (
	ainame string
	url string
	port string
	config []byte
	fetch []string = []string{"name","score"}
)

func init() {
	s := g.Server()

	//url全小写
	s.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)

	//AI组 post{"image":base64}(config) return {"name":"","score":""}(fetch)
	s.Group("/ai", func(group *ghttp.RouterGroup) {
		name := (&ai.Config{
			Config: config,
			Url:    url,
			Port:   port,
			Fetch:  fetch,
		}).Client()
		s.BindObject("/{.struct}/", name)
	})

}
