package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gins"
	"gitee.com/johng/gf/g/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})
	s.SetPort(8100)
	s.SetServerRoot("/Users/zhangyong/Movies/美拍")
	gins.View().AddPath(".")
	s.BindHandler("/template2", func(r *ghttp.Request) {
		content, _ := gins.View().Parse("index.tpl", map[string]interface{}{
			"id":   123,
			"name": "john",
		})
		r.Response.Write(content)
	})
	s.Run()
}
