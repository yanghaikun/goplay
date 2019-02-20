package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	_ "net/http/pprof"
	"time"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数, 默认是不会解析的
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello world") //输出到客户端的信息

	for i := 0; i < 100000; i++  {
		time.Sleep(time.Microsecond)
	}
}

func ServeHttp() {
	log.Printf("run hello world server")
	http.HandleFunc("/", greeting)           //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	ServeHttp()
}
