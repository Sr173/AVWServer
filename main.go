package main

import (
	"./protocol"
	"net/http"
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

	switch r.Method {
	case http.MethodGet:
		protocol.HttpGetHandler(w,r)
		break
	case http.MethodPost:
		protocol.HttpPostHandler(w,r)
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
