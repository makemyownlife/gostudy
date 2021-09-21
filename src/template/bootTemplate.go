package template

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Inventory2 struct {
	GroupId          string
	ProjectName      string
	ModuleNamePrefix string
	BasePackage      string
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
		basePackage := config.Get("maven.basePackage").(string)

		log.Println("项目根路径：" + path)
		log.Println("项目名:" + projectName)
		log.Println("文件分隔符:" + string(os.PathSeparator))
		log.Println("模块名前缀：" + moduleNamePrefix)
		log.Println("base包目录：" + basePackage)

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
		pomObject := Inventory2{groupId, projectName, moduleNamePrefix, basePackage}
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

		var separator = string(os.PathSeparator)

		//复制 pom gitignore文件到目的目录
		ignoreFilePath := []string{projectPath, ".gitignore"}
		copyStaticFileToTarget("config/boot/static/.gitignore", strings.Join(ignoreFilePath, string(os.PathSeparator)))

		//===================================================================================================创建common 模块 ===================================================================================================
		var commonModule = moduleNamePrefix + "-common"
		var commonPath = projectPath + string(os.PathSeparator) + commonModule
		os.MkdirAll(commonPath, 0777)

		var commonSrcPath = commonPath + string(os.PathSeparator) + "src"
		log.Println("commonSrcPath:   " + commonSrcPath)
		createSrcDir(commonSrcPath)

		arr := strings.Split(basePackage, ".")
		packageStr := strings.Join(arr, string(os.PathSeparator))
		log.Println("包目录：    " + packageStr)

		var commonMainPath = commonSrcPath + separator + "main"
		var commonJavaPath = commonMainPath + separator + "java"
		var commonResPath = commonMainPath + separator + "resources"
		var commonTestPath = commonSrcPath + separator + "test"
		var commonTestJavaPath = commonTestPath + separator + "java"

		createSrcDir(commonJavaPath)
		createSrcDir(commonResPath)
		createSrcDir(commonTestPath)
		createSrcDir(commonTestJavaPath)

		commonClasspath := commonJavaPath + string(os.PathSeparator) + packageStr + string(os.PathSeparator) + "common"
		createSrcDir(commonClasspath)

		var commonPackageInfo = "package " + basePackage + ".common;"
		ioutil.WriteFile(commonClasspath+separator+"package-info.java", []byte(string(commonPackageInfo)), 0777)

		renderPomFile("config/boot/static/common/pom.xml", &pomObject, commonPath)

		//===================================================================================================创建api 模块 ===================================================================================================
		var apiModule = moduleNamePrefix + "-api"
		var apiPath = projectPath + string(os.PathSeparator) + apiModule
		os.MkdirAll(apiPath, 0777)

		var apiSrcPath = apiPath + string(os.PathSeparator) + "src"
		log.Println("apiSrcPath:   " + apiSrcPath)
		createSrcDir(apiSrcPath)

		var apiMainPath = apiSrcPath + separator + "main"
		var apiJavaPath = apiMainPath + separator + "java"
		var apiResPath = apiMainPath + separator + "resources"
		var apiTestPath = apiSrcPath + separator + "test"
		var apiTestJavaPath = apiTestPath + separator + "java"

		createSrcDir(apiJavaPath)
		createSrcDir(apiResPath)
		createSrcDir(apiTestPath)
		createSrcDir(apiTestJavaPath)

		apiClasspath := apiJavaPath + string(os.PathSeparator) + packageStr + string(os.PathSeparator) + "api"
		createSrcDir(apiClasspath)

		var apiPackageInfo = "package " + basePackage + ".api;"
		ioutil.WriteFile(apiClasspath+separator+"package-info.java", []byte(string(apiPackageInfo)), 0777)

		renderPomFile("config/boot/static/api/pom.xml", &pomObject, apiPath)

		//===================================================================================================创建domain模块 ===================================================================================================
		var domainModule = moduleNamePrefix + "-domain"
		var domainPath = projectPath + string(os.PathSeparator) + domainModule
		os.MkdirAll(domainPath, 0777)

		var domainSrcPath = domainPath + string(os.PathSeparator) + "src"
		log.Println("domainSrcPath:   " + domainSrcPath)
		createSrcDir(domainSrcPath)

		var domainMainPath = domainSrcPath + separator + "main"
		var domainJavaPath = domainMainPath + separator + "java"
		var domainResPath = domainMainPath + separator + "resources"
		var domainTestPath = domainSrcPath + separator + "test"
		var domainTestJavaPath = domainTestPath + separator + "java"

		createSrcDir(domainJavaPath)
		createSrcDir(domainResPath)
		createSrcDir(domainTestPath)
		createSrcDir(domainTestJavaPath)

		domainClasspath := domainJavaPath + string(os.PathSeparator) + packageStr + string(os.PathSeparator) + "domain"
		createSrcDir(domainClasspath)

		var domainPackageInfo = "package " + basePackage + ".domain;"
		ioutil.WriteFile(domainClasspath+separator+"package-info.java", []byte(string(domainPackageInfo)), 0777)

		renderPomFile("config/boot/static/domain/pom.xml", &pomObject, domainPath)
		createSrcDir(domainClasspath + separator + "mapper")

		renderOtherFile(
			"config/boot/static/domain/UserMapper.java",
			&pomObject,
			domainClasspath+separator+"mapper"+separator+"UserMapper.java")

		createSrcDir(domainClasspath + separator + "po")

		renderOtherFile(
			"config/boot/static/domain/User.java",
			&pomObject,
			domainClasspath+separator+"po"+separator+"User.java")

		//===================================================================================================创建service模块 ===================================================================================================

		var serviceModule = moduleNamePrefix + "-service"
		var servicePath = projectPath + string(os.PathSeparator) + serviceModule
		os.MkdirAll(servicePath, 0777)

		var serviceSrcPath = servicePath + string(os.PathSeparator) + "src"
		log.Println("serviceSrcPath:   " + serviceSrcPath)
		createSrcDir(serviceSrcPath)

		var serviceMainPath = serviceSrcPath + separator + "main"
		var serviceJavaPath = serviceMainPath + separator + "java"
		var serviceResPath = serviceMainPath + separator + "resources"
		var serviceTestPath = serviceSrcPath + separator + "test"
		var serviceTestJavaPath = serviceTestPath + separator + "java"

		createSrcDir(serviceJavaPath)
		createSrcDir(serviceResPath)
		createSrcDir(serviceTestPath)
		createSrcDir(serviceTestJavaPath)

		serviceClasspath := serviceJavaPath + string(os.PathSeparator) + packageStr + string(os.PathSeparator) + "service"
		createSrcDir(serviceClasspath)

		var servicePackageInfo = "package " + basePackage + ".service;"
		ioutil.WriteFile(serviceClasspath+separator+"package-info.java", []byte(string(servicePackageInfo)), 0777)

		renderPomFile("config/boot/static/service/pom.xml", &pomObject, servicePath)
		//===================================================================================================创建provider模块 ===================================================================================================
		var providerModule = moduleNamePrefix + "-provider"
		var providerPath = projectPath + string(os.PathSeparator) + providerModule
		os.MkdirAll(providerPath, 0777)

		var providerSrcPath = providerPath + string(os.PathSeparator) + "src"
		log.Println("providerSrcPath:   " + providerSrcPath)
		createSrcDir(providerSrcPath)

		var providerMainPath = providerSrcPath + separator + "main"
		var providerJavaPath = providerMainPath + separator + "java"
		var providerResPath = providerMainPath + separator + "resources"
		var providerTestPath = providerSrcPath + separator + "test"
		var providerTestJavaPath = providerTestPath + separator + "java"

		createSrcDir(providerJavaPath)
		createSrcDir(providerResPath)
		createSrcDir(providerTestPath)
		createSrcDir(providerTestJavaPath)

		providerClasspath := providerJavaPath + string(os.PathSeparator) + packageStr + string(os.PathSeparator) + "provider"
		createSrcDir(providerClasspath)

		var providerPackageInfo = "package " + basePackage + ".provider;"
		ioutil.WriteFile(providerClasspath+separator+"package-info.java", []byte(string(providerPackageInfo)), 0777)

		renderPomFile("config/boot/static/provider/pom.xml", &pomObject, providerPath)

		//===================================================================================================创建server模块 ===================================================================================================
		var serverModule = moduleNamePrefix + "-server"
		var serverPath = projectPath + string(os.PathSeparator) + serverModule
		os.MkdirAll(serverPath, 0777)

		var serverSrcPath = serverPath + string(os.PathSeparator) + "src"
		log.Println("serverSrcPath:   " + serverSrcPath)
		createSrcDir(serverSrcPath)

		var serverMainPath = serverSrcPath + separator + "main"
		var serverJavaPath = serverMainPath + separator + "java"
		var serverResPath = serverMainPath + separator + "resources"
		var serverTestPath = serverSrcPath + separator + "test"
		var serverTestJavaPath = serverTestPath + separator + "java"

		createSrcDir(serverJavaPath)
		createSrcDir(serverResPath)
		createSrcDir(serverTestPath)
		createSrcDir(serverTestJavaPath)

		serverClasspath := serverJavaPath + string(os.PathSeparator) + packageStr + string(os.PathSeparator) + "server"
		createSrcDir(serverClasspath)

		var serverPackageInfo = "package " + basePackage + ".server;"
		ioutil.WriteFile(serverClasspath+separator+"package-info.java", []byte(string(serverPackageInfo)), 0777)

		renderPomFile("config/boot/static/server/pom.xml", &pomObject, serverPath)
		//===================================================================================================创建demo模块 ===================================================================================================

	}

}
