package ai

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

//read serving_client_conf.stream.prototxt as byte
//target server url and port
//what u fetch
type Config struct{
	Config []byte
	Url string
	Port string
	Fetch []string
}

type MyHandle struct {
	Handle Handle
	Fetch []string
}

//set config of model client
func (c *Config) Client() *MyHandle {
	handle := LoadModelConfig(c.Config)

	return &MyHandle{
		Connect(c.Url, c.Port, handle),
		c.Fetch,
	}
}

//input k-(byte) predict what u fetch return json
func (h *MyHandle) Predict(r *ghttp.Request){
	var data map[string][]byte
	if ok := r.Parse(&data); ok != nil {
		glog.Debug(ok)
	}
	r.Response.WriteJsonExit(
	Predict(h.Handle, data, h.Fetch ),
	)
}