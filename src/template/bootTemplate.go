package template

import (
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

func CreateBootMavenProject() {
	execDirAbsPath, _ := os.Getwd()
	log.Println("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)
	config, err := toml.LoadFile(execDirAbsPath + "/config/boot/config.toml")
	if err != nil {
		log.Println("Error ", err.Error())
	} else {
		path := config.Get("maven.path").(string)
		projectName := config.Get("maven.projectName").(string)
		moduleNamePrefix := config.Get("maven.moduleNamePrefix").(string)

		log.Println("项目根路径：" + path)
		log.Println("项目名:" + projectName)
		log.Println("文件分隔符:" + string(os.PathSeparator))
		log.Println("模块名前缀：" + moduleNamePrefix)

		projectPath := path + string(os.PathSeparator) + projectName
		log.Println("trying to create ProjectPath:    " + projectPath)

		//开始创建工程目录
		err := os.MkdirAll(projectPath, 0777)
		if err != nil {
			log.Printf("%s", err)
			return
		}
		log.Println("Create Direcry OK!")

	}
}
