package main

import (
	"AVWServer/protocol"
	"net/http"
	"strings"
)

var func_map map[string]map[string]http.Handler


func scho(w http.ResponseWriter, r*http.Request) {
	w.Write([]byte("请求就绪"))
}

func acho(w http.ResponseWriter, r*http.Request) {
	w.Write([]byte("api"))
}

func dispatch(w http.ResponseWriter, r *http.Request){
	if r.RequestURI == "/ws"{
		protocol.WeosocketHandler(w,r)
		return
	}
	getPath_ := strings.Split(r.RequestURI,"/")
	println(r.Method,r.RequestURI,getPath_[1])

	switch r.Method {
	case http.MethodGet:
		break
	case http.MethodPost:
		break
	case http.MethodDelete:
		break
	case http.MethodConnect:
		break
	}
}

func main() {
	http.HandleFunc("/", dispatch)
	http.ListenAndServe("localhost:80", nil)
}
