package main

import (
	"github.com/pelletier/go-toml"
	"go-study/src/template"
	"log"
	"os"
)

func main() {
	execDirAbsPath, _ := os.Getwd()
	log.Println("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)
	config, err := toml.LoadFile(execDirAbsPath + "/config/boot/config.toml")
	if err != nil {
		log.Println("Error ", err.Error())
	} else {
		isDubbo := config.Get("maven.isDubbo").(string)
		log.Println("是否dubbo项目：" + isDubbo)
		if isDubbo == "1" {
			template.CreateDubboMavenProject()
		} else {
			template.CreateBootMavenProject()
		}
	}
	//执行 MyBatis generator 命令

}
