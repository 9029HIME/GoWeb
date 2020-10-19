package main

import (
	"fmt"
	"net/http"
)

func fakeMain() {
	mux:= http.NewServeMux();
	mux.HandleFunc("/",handler)
	http.ListenAndServe(":8082",mux)
}

func handler(response http.ResponseWriter,request *http.Request){
	fmt.Fprintln(response,"HelloWorld!",request.URL.Path)
	response.Write([]byte("自定义多路复用器"))
}
