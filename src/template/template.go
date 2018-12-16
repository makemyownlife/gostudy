package template

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Inventory struct {
	GroupId    string
	ArtifactId string
}

func CreateMavenProject() {
	//从静态资源中获取数据
	execDirAbsPath, _ := os.Getwd()
	log.Println("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)
	config, err := toml.LoadFile(execDirAbsPath + "/config/template/config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		path := config.Get("maven.path").(string)
		projectName := config.Get("maven.projectName").(string)
		fileSeperator := config.Get("maven.fileSeperator").(string)
		groupId := config.Get("maven.groupId").(string)
		artifactId := config.Get("maven.artifactId").(string)

		fmt.Println("项目根路径：" + path)
		fmt.Println("项目名:" + projectName)
		fmt.Println("文件分割符号：" + fileSeperator)

		projectPath := path + projectName
		fmt.Println("trying to create ProjectPath:" + projectPath)

		//开始创建工程目录
		err := os.MkdirAll(projectPath, 0777)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		fmt.Println("Create Direcry OK!")

		//创建pom文件
		pomObject := Inventory{groupId, artifactId}
		//创建pom文件目录
		pomStr := readStaticFile("config/template/static/pom.xml")
		pomTmp, err := template.New("pom").Parse(pomStr) //建立一个模板
		//将struct与模板合成，合成结果放到os.Stdout里
		pomWriter, err3 := os.Create(projectPath + fileSeperator + "pom.xml") //创建文件
		err = pomTmp.Execute(pomWriter, pomObject)
		if err != nil {
			panic(err)
		}
		if err3 != nil {
			panic(err)
		}
		defer pomWriter.Close()
		fmt.Println("create pomfile success")

		//复制 pom gitignore文件到目的目录
		ignoreFilePath := []string{projectPath, ".gitignore"}
		copyStaticFileToTarget("config/template/static/.gitignore", strings.Join(ignoreFilePath, fileSeperator))

		//循环创建 src 相关的目录 以及log4j.xml
		javaPath := projectPath + fileSeperator + "src" + fileSeperator + "main" + fileSeperator + "java"
		createSrcDir(javaPath)

		//将groupId对应的字符串组成目录形式
		arr := strings.Split(groupId, ".")
		packageStr := strings.Join(arr, fileSeperator)
		createSrcDir(javaPath + fileSeperator + packageStr)

		resourcesPath := projectPath + fileSeperator + "src" + fileSeperator + "main" + fileSeperator + "resources"
		createSrcDir(resourcesPath)
		testPath := projectPath + fileSeperator + "src" + fileSeperator + "test" + fileSeperator + "java"
		createSrcDir(testPath)
		testResourcePath := projectPath + fileSeperator + "src" + fileSeperator + "test" + fileSeperator + "resources"
		createSrcDir(testResourcePath)
		createSrcDir(testPath + fileSeperator + packageStr)

		//logback页面的修正
		logbackStr := readStaticFile("config/template/static/logback.xml")
		ioutil.WriteFile(resourcesPath+fileSeperator+"logback.xml", []byte(logbackStr), 0777)
	}

}
