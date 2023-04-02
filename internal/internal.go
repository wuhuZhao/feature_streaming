package internal

import (
	"encoding/json"
	"github.com/wuhuZhao/feature_streaming/internal/consumer"
	"github.com/wuhuZhao/feature_streaming/internal/execute"
	"github.com/wuhuZhao/feature_streaming/internal/producer"

	"fmt"
	"io"
	"net/http"
)

var exec execute.Execute

var queue = make(chan []byte, 100)

func init() {
	exec = execute.NewDefaultExecute(producer.NewKafkaProducer("127.0.0.1:8888", "zhk", 0, queue), consumer.NewDefaultConsumer(queue))
	if err := exec.Init(); err != nil {
		panic(err)
	}
}

func InternalHandler() func(response http.ResponseWriter, r *http.Request) {
	var handler func(response http.ResponseWriter, r *http.Request)
	handler = func(response http.ResponseWriter, request *http.Request) {
		data, err := io.ReadAll(request.Body)
		if err != nil {
			response.Write([]byte(fmt.Sprintf("err: %v", err)))
			return
		}
		dataMap := map[string]string{}
		if err := json.Unmarshal(data, &dataMap); err != nil {
			response.Write([]byte(fmt.Sprintf("err: %v", err)))
			return
		}
		config := dataMap["config"]
		if len(config) == 0 {
			response.Write([]byte("err: config is null"))
			return
		}
		exec.Exec(config)
		response.Write([]byte("start task!"))
	}
	return handler
}
