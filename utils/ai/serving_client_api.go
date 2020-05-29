package ai

import (
	"encoding/json"
	pb "github.com/PaddlePaddle/Serving/go/proto"
	"github.com/gogf/gf/net/ghttp"
	"github.com/golang/protobuf/proto"
	"log"
)

type Tensor struct {
	Data   []byte `json:"data"`
	FloatData	   []float32 `json:"float_data"`
	IntData	   []int `json:"int_data"`
	Int64Data	   []int64 `json:"int64_data"`
	ElemType	int `json:"elem_type"`
	Shape	[]int `json:"shape"`
}

type FeedInst struct {
	TensorArray     []Tensor `json:"tensor_array"`
}

type FetchInst struct {
	TensorArray      []Tensor `json:"tensor_array"`
}

type Request struct {
	Insts   []FeedInst `json:"insts"`
	FetchVarNames	[]string `json:"fetch_var_names"`
	ProfileServer	bool `json:"profile_server"`
}

type Response struct {
	Insts    []FetchInst `json:"insts"`
	ProfileTime	  []int64 `json:"profile_time"`
}

type Handle struct {
	Url    string
	Port   string
	FeedAliasNameMap	map[string]string
	FeedShapeMap	map[string][]int
	FeedNameMap   map[string]int
	FeedAliasNames	   []string
	FetchNameMap  map[string]int
	FetchAliasNameMap	map[string]string
}

//read config as []byte
func LoadModelConfig(config []byte) Handle {
	general_model_config := &pb.GeneralModelConfig{}
	if err := proto.Unmarshal(config, general_model_config); err != nil {
		log.Fatalln("Failed to parse GeneralModelConfig: ", err)
	}
	log.Println("read protobuf succeed")
	handle := Handle{}
	handle.FeedNameMap = map[string]int{}
	handle.FeedAliasNameMap = map[string]string{}
	handle.FeedShapeMap = map[string][]int{}
	handle.FetchNameMap = map[string]int{}
	handle.FetchAliasNameMap = map[string]string{}
	handle.FeedAliasNames = []string{}

	for i, v := range general_model_config.FeedVar {
		handle.FeedNameMap[*v.Name] = i
		tmp_array := []int{}
		for _, vv := range v.Shape {
			tmp_array = append(tmp_array, int(vv))
		}
		handle.FeedShapeMap[*v.Name] = tmp_array
		handle.FeedAliasNameMap[*v.AliasName] = *v.Name
		handle.FeedAliasNames = append(handle.FeedAliasNames, *v.AliasName)
	}

	for i, v := range general_model_config.FetchVar {
		handle.FetchNameMap[*v.Name] = i
		handle.FetchAliasNameMap[*v.AliasName] = *v.Name
	}

	return handle
}

func Connect(url string, port string, handle Handle) Handle {
	handle.Url = url
	handle.Port = port
	return handle
}

func Predict(handle Handle, byte_feed_map map[string][]byte, fetch []string) []byte {


	var tensor_array []Tensor
	var inst FeedInst
	tensor_array = []Tensor{}
	inst = FeedInst{}

	for i := 0; i < len(handle.FeedAliasNames); i++ {
		key_i := handle.FeedAliasNames[i]
		var tmp Tensor
		tmp.IntData = []int{}
		tmp.Int64Data = []int64{}
		tmp.Shape = []int{}
		tmp.Data = byte_feed_map[key_i]
		tmp.ElemType = 0
		tmp.Shape = handle.FeedShapeMap[key_i]
		tensor_array = append(tensor_array, tmp)
	}

	inst.TensorArray = tensor_array

	var profile_server bool
	profile_server = false

	req := &Request{
		Insts: []FeedInst{inst},
		FetchVarNames: fetch,
		ProfileServer: profile_server}

	b, _ := json.Marshal(req)

	resp := ghttp.PostBytes(
		"http://" + handle.Url + ":" + handle.Port + "/",
		b,
		)

	return resp

}
