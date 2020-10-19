package main

import (
	"fmt"
	"net/http"
)

func main() {
		//TODO 会将传入的handler函数转换为ServeHTTP()函数，前提是参数声明必须一致
		http.HandleFunc("/",handler)
		f := new(fuck)
		//也可以使用自定义的Handler(实现ServeHTTP接口的类)
		http.Handle("/fuck",f)
		//nil:用默认的多路复用器 DefaultServeMux
		http.ListenAndServe(":8080",nil)
}

//创建处理器 TODO 处理器的参数顺序不能变
func handler(response http.ResponseWriter,request *http.Request){
	fmt.Fprintln(response,"HelloWorld!",request.URL.Path)
	response.Write([]byte("你好世界"))
}

type fuck struct{

}

func (f *fuck) ServeHTTP(response http.ResponseWriter,request *http.Request){
	response.Write([]byte("fuckyou"))
}