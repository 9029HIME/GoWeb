package main

import (
	"net/http"
	"time"
)

func main() {
	myh := new(MyH)
	//TODO 除了用http.ListenAndServe以外，还能自己自定义一个服务器
	server:= http.Server{
		Addr: ":8081",// 监听8081端口
		Handler: myh,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}

type MyH struct {

}

func (my *MyH) ServeHTTP(response http.ResponseWriter,request *http.Request){
	response.Write([]byte("自定义的"))
}


