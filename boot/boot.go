package boot

import (
	"gf-app/utils/ai"
	"gf-app/utils/configcli"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/os/glog"
	"github.com/micro/go-micro/v2/config"
)

var Configs []ai.Config

func syncconfig() {
	c, err := config.NewConfig(
		config.WithSource(
			configcli.NewSource("app", "cluster", "ai", configcli.WithURL("http://127.0.0.1:8080/"))))
	if err != nil {
		panic(err)
	}
	glog.Info("read: ", string(c.Get().Bytes()))

	// Watch 返回前 micro config 会调用 Read 读一次配置
	w, err := c.Watch()
	if err != nil {
		panic(err)
	}

	for {
		// 会比较 value，内容不变不会返回
		v, err := w.Next()
		if err != nil {
			panic(err)
		}

		gcfg.SetContent(string(v.Bytes()))
		_ = ghttp.RestartAllServer()
	}
}

func init()  {
	go syncconfig()
}