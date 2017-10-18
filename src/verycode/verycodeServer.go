package verycode

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func StartVeryCodeServer() {
	http.HandleFunc("/", getVeryCodeNum)     //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getVeryCodeNum(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数, 默认是不会解析的
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path:", r.URL.Path)
	s := strings.Join(r.Form["s"], "");
	//输出到客户端的信息
	fmt.Fprintf(w, s)
}
