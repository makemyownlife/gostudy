package template

import (
	"github.com/pelletier/go-toml"
	"log"
	"os"
	"text/template"
)

type Inventory2 struct {
	GroupId          string
	ProjectName      string
	ModuleNamePrefix string
}

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
		groupId := config.Get("maven.groupId").(string)

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

		//创建pom文件
		pomObject := Inventory2{groupId, projectName, moduleNamePrefix}
		//创建pom文件目录
		pomStr := readStaticFile("config/boot/static/pom.xml")
		pomTmp, err := template.New("pom").Parse(pomStr) //建立一个模板
		//将struct与模板合成，合成结果放到os.Stdout里
		var pomPath = projectPath + string(os.PathSeparator) + "pom.xml"
		log.Println("create pomfile start pomPath:" + pomPath)
		pomWriter, err3 := os.Create(pomPath) //创建文件
		err = pomTmp.Execute(pomWriter, pomObject)
		if err != nil {
			panic(err)
		}
		if err3 != nil {
			panic(err)
		}
		defer pomWriter.Close()
		log.Println("create pomfile success")
	}
}
