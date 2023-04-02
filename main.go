package main

import (
	"github.com/wuhuZhao/feature_streaming/internal"
	"net/http"
)

func main() {
	http.HandleFunc("/start", internal.InternalHandler())
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		panic(err)
	}
}
