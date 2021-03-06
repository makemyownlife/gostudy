package main

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/frame/gins"
	"github.com/gogf/gf/g/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})
	s.SetPort(8100)
	gins.View().AddPath(".")
	//本地文件系统
	//s.SetServerRoot("D:/AdminLTE-master")
	s.BindHandler("/template2", func(r *ghttp.Request) {
		content, _ := gins.View().Parse("index.tpl", map[string]interface{}{
			"id":   123,
			"name": "john",
		})
		r.Response.Write(content)
	})
	s.Run()
}
