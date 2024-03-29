package template

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Inventory2 struct {
	GroupId          string
	ProjectName      string
	ModuleNamePrefix string
	BasePackage      string
}

func CreateDubboMavenProject() {
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
		log.Println("文件分隔符:" + string(filepath.Separator))
		log.Println("模块名前缀：" + moduleNamePrefix)
		log.Println("base包目录：" + basePackage)

		projectPath := path + string(filepath.Separator) + projectName
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
		pomStr := readStaticFile("config/boot/static/pom-dubbo.xml")
		pomTmp, err := template.New("pom").Parse(pomStr) //建立一个模板
		//将struct与模板合成，合成结果放到os.Stdout里
		var pomPath = projectPath + string(filepath.Separator) + "pom.xml"
		log.Println("create pomfile start pomPath:" + pomPath)
		pomWriter, err3 := os.Create(pomPath) //创建文件
		err = pomTmp.Execute(pomWriter, pomObject)
		if err != nil {
			panic(err)
		}
		if err3 != nil {
			panic(err)
		}
		//	defer pomWriter.Close()
		log.Println("create pomfile success")

		var separator = string(filepath.Separator)

		//复制 pom gitignore文件到目的目录
		ignoreFilePath := []string{projectPath, ".gitignore"}
		copyStaticFileToTarget("config/boot/static/.gitignore", strings.Join(ignoreFilePath, string(filepath.Separator)))

		//===================================================================================================创建common 模块 ===================================================================================================
		var commonModule = moduleNamePrefix + "-common"
		var commonPath = projectPath + string(filepath.Separator) + commonModule
		os.MkdirAll(commonPath, 0777)

		var commonSrcPath = commonPath + string(filepath.Separator) + "src"
		log.Println("commonSrcPath:   " + commonSrcPath)
		createSrcDir(commonSrcPath)

		arr := strings.Split(basePackage, ".")
		packageStr := strings.Join(arr, string(filepath.Separator))
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

		commonClasspath := commonJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "common"
		createSrcDir(commonClasspath)

		var commonPackageInfo = "package " + basePackage + ".common;"
		ioutil.WriteFile(commonClasspath+separator+"package-info.java", []byte(string(commonPackageInfo)), 0777)

		renderPomFile("config/boot/static/common/pom.xml", &pomObject, commonPath)

		//===================================================================================================创建api 模块 ===================================================================================================
		var apiModule = moduleNamePrefix + "-api"
		var apiPath = projectPath + string(filepath.Separator) + apiModule
		os.MkdirAll(apiPath, 0777)

		var apiSrcPath = apiPath + string(filepath.Separator) + "src"
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

		apiClasspath := apiJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "api"
		createSrcDir(apiClasspath)

		//var apiPackageInfo = "package " + basePackage + ".api;"
		//ioutil.WriteFile(apiClasspath+separator+"package-info.java", []byte(string(apiPackageInfo)), 0777)

		renderPomFile("config/boot/static/api/pom.xml", &pomObject, apiPath)

		createSrcDir(apiClasspath + separator + "dto")

		renderOtherFile(
			"config/boot/static/api/dto/DubboRpcResult.java",
			&pomObject,
			apiClasspath+separator+"dto"+separator+"DubboRpcResult.java")

		renderOtherFile(
			"config/boot/static/api/TestDubboService.java",
			&pomObject,
			apiClasspath+separator+"TestDubboService.java")

		//===================================================================================================创建domain模块 ===================================================================================================
		var domainModule = moduleNamePrefix + "-domain"
		var domainPath = projectPath + string(filepath.Separator) + domainModule
		os.MkdirAll(domainPath, 0777)

		var domainSrcPath = domainPath + string(filepath.Separator) + "src"
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

		domainClasspath := domainJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "domain"
		createSrcDir(domainClasspath)

		var domainPackageInfo = "package " + basePackage + ".domain;"
		ioutil.WriteFile(domainClasspath+separator+"package-info.java", []byte(string(domainPackageInfo)), 0777)

		renderPomFile("config/boot/static/domain/pom.xml", &pomObject, domainPath)
		createSrcDir(domainClasspath + separator + "mapper")

		renderOtherFile(
			"config/boot/static/domain/TestMapper.java",
			&pomObject,
			domainClasspath+separator+"mapper"+separator+"TestMapper.java")

		createSrcDir(domainClasspath + separator + "po")

		renderOtherFile(
			"config/boot/static/domain/TestPo.java",
			&pomObject,
			domainClasspath+separator+"po"+separator+"TestPo.java")

		mapperDir := domainResPath + separator + "mapper"
		createSrcDir(mapperDir)

		renderOtherFile(
			"config/boot/static/domain/TestMapper.xml",
			&pomObject,
			mapperDir+separator+"TestMapper.xml")

		//===================================================================================================创建service模块 ===================================================================================================

		var serviceModule = moduleNamePrefix + "-service"
		var servicePath = projectPath + string(filepath.Separator) + serviceModule
		os.MkdirAll(servicePath, 0777)

		var serviceSrcPath = servicePath + string(filepath.Separator) + "src"
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

		serviceClasspath := serviceJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "service"
		createSrcDir(serviceClasspath)

		renderPomFile("config/boot/static/service/pom.xml", &pomObject, servicePath)

		renderOtherFile(
			"config/boot/static/service/TestService.java",
			&pomObject,
			serviceClasspath+separator+"TestService.java")

		//===================================================================================================创建provider模块 ===================================================================================================
		var providerModule = moduleNamePrefix + "-provider"
		var providerPath = projectPath + string(filepath.Separator) + providerModule
		os.MkdirAll(providerPath, 0777)

		var providerSrcPath = providerPath + string(filepath.Separator) + "src"
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

		providerClasspath := providerJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "provider"
		createSrcDir(providerClasspath)

		renderOtherFile(
			"config/boot/static/provider/TestDubboServiceImpl.java",
			&pomObject,
			providerClasspath+separator+"TestDubboServiceImpl.java")

		renderPomFile("config/boot/static/provider/pom.xml", &pomObject, providerPath)

		//===================================================================================================创建server模块 ===================================================================================================
		var serverModule = moduleNamePrefix + "-server"
		var serverPath = projectPath + string(filepath.Separator) + serverModule
		os.MkdirAll(serverPath, 0777)

		var serverSrcPath = serverPath + string(filepath.Separator) + "src"
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

		serverClasspath := serverJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "server"
		createSrcDir(serverClasspath)

		createSrcDir(serverClasspath + string(filepath.Separator) + "config")
		renderOtherFile(
			"config/boot/static/server-boot/config/RedisConfig.java",
			&pomObject,
			serverClasspath+separator+"config"+separator+"RedisConfig.java")
		renderOtherFile(
			"config/boot/static/server-boot/config/CorsConfig.java",
			&pomObject,
			serverClasspath+separator+"config"+separator+"CorsConfig.java")
		renderOtherFile(
			"config/boot/static/server-boot/config/WebMvcConfig.java",
			&pomObject,
			serverClasspath+separator+"config"+separator+"WebMvcConfig.java")

		renderPomFile("config/boot/static/server/pom.xml", &pomObject, serverPath)

		renderOtherFile(
			"config/boot/static/server/MainApplication.java",
			&pomObject,
			serverClasspath+separator+"MainApplication.java")

		renderOtherFile(
			"config/boot/static/server/application.yml",
			&pomObject,
			serverResPath+separator+"application.yml")

		renderOtherFile(
			"config/boot/static/server/application-prod.yml",
			&pomObject,
			serverResPath+separator+"application-prod.yml")

		renderOtherFile(
			"config/boot/static/server/application-test.yml",
			&pomObject,
			serverResPath+separator+"application-test.yml")

		createSrcDir(serverResPath + separator + "logger")

		renderOtherFile("config/boot/static/server/logback-prod.xml", &pomObject, serverResPath+separator+"logger"+separator+"logback-prod.xml")
		renderOtherFile("config/boot/static/server/logback-test.xml", &pomObject, serverResPath+separator+"logger"+separator+"logback-test.xml")

		// 添加打包相关内容
		createSrcDir(serverMainPath + separator + "assembly")
		renderOtherFile("config/boot/static/server/assembly/assembly.xml", &pomObject, serverMainPath+separator+"assembly"+separator+"assembly.xml")
		createSrcDir(serverMainPath + separator + "assembly" + separator + "bin")
		renderOtherFile("config/boot/static/server/assembly/binTemp/start.sh.temp", &pomObject, serverMainPath+separator+"assembly"+separator+"bin"+separator+"start.sh")
		renderOtherFile("config/boot/static/server/assembly/binTemp/stop.sh.temp", &pomObject, serverMainPath+separator+"assembly"+separator+"bin"+separator+"stop.sh")

		//===================================================================================================创建demo模块 ===================================================================================================

	}

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
		log.Println("文件分隔符:" + string(filepath.Separator))
		log.Println("模块名前缀：" + moduleNamePrefix)
		log.Println("base包目录：" + basePackage)

		projectPath := path + string(filepath.Separator) + projectName
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
		var pomPath = projectPath + string(filepath.Separator) + "pom.xml"
		log.Println("create pomfile start pomPath:" + pomPath)
		pomWriter, err3 := os.Create(pomPath) //创建文件
		err = pomTmp.Execute(pomWriter, pomObject)
		if err != nil {
			panic(err)
		}
		if err3 != nil {
			panic(err)
		}
		//defer pomWriter.Close()
		log.Println("create pomfile success")

		var separator = string(filepath.Separator)

		//复制 pom gitignore文件到目的目录
		ignoreFilePath := []string{projectPath, ".gitignore"}
		copyStaticFileToTarget("config/boot/static/.gitignore", strings.Join(ignoreFilePath, string(filepath.Separator)))

		//===================================================================================================创建common 模块 ===================================================================================================
		var commonModule = moduleNamePrefix + "-common"
		var commonPath = projectPath + string(filepath.Separator) + commonModule
		os.MkdirAll(commonPath, 0777)

		var commonSrcPath = commonPath + string(filepath.Separator) + "src"
		log.Println("commonSrcPath:   " + commonSrcPath)
		createSrcDir(commonSrcPath)

		arr := strings.Split(basePackage, ".")
		packageStr := strings.Join(arr, string(filepath.Separator))
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

		commonClasspath := commonJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "common"
		createSrcDir(commonClasspath)

		var commonPackageInfo = "package " + basePackage + ".common;"
		ioutil.WriteFile(commonClasspath+separator+"package-info.java", []byte(string(commonPackageInfo)), 0777)

		createSrcDir(commonClasspath + string(filepath.Separator) + "result")

		renderOtherFile(
			"config/boot/static/common/result/ResponseEntity.java",
			&pomObject,
			commonClasspath+separator+"result"+separator+"ResponseEntity.java")

		renderOtherFile(
			"config/boot/static/common/result/ResultPage.java",
			&pomObject,
			commonClasspath+separator+"result"+separator+"ResultPage.java")

		renderOtherFile(
			"config/boot/static/common/result/PageDataInfo.java",
			&pomObject,
			commonClasspath+separator+"result"+separator+"PageDataInfo.java")

		renderPomFile("config/boot/static/common/pom.xml", &pomObject, commonPath)

		//===================================================================================================创建domain模块 ===================================================================================================
		var domainModule = moduleNamePrefix + "-domain"
		var domainPath = projectPath + string(filepath.Separator) + domainModule
		os.MkdirAll(domainPath, 0777)

		var domainSrcPath = domainPath + string(filepath.Separator) + "src"
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

		domainClasspath := domainJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "domain"
		createSrcDir(domainClasspath)

		var domainPackageInfo = "package " + basePackage + ".domain;"
		ioutil.WriteFile(domainClasspath+separator+"package-info.java", []byte(string(domainPackageInfo)), 0777)

		renderPomFile("config/boot/static/domain/pom.xml", &pomObject, domainPath)
		createSrcDir(domainClasspath + separator + "mapper")

		renderOtherFile(
			"config/boot/static/domain/TestMapper.java",
			&pomObject,
			domainClasspath+separator+"mapper"+separator+"TestMapper.java")

		createSrcDir(domainClasspath + separator + "po")

		renderOtherFile(
			"config/boot/static/domain/TestPo.java",
			&pomObject,
			domainClasspath+separator+"po"+separator+"TestPo.java")

		mapperDir := domainResPath + separator + "mapper"
		createSrcDir(mapperDir)

		renderOtherFile(
			"config/boot/static/domain/TestMapper.xml",
			&pomObject,
			mapperDir+separator+"TestMapper.xml")

		//===================================================================================================创建service模块 ===================================================================================================

		var serviceModule = moduleNamePrefix + "-service"
		var servicePath = projectPath + string(filepath.Separator) + serviceModule
		os.MkdirAll(servicePath, 0777)

		var serviceSrcPath = servicePath + string(filepath.Separator) + "src"
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

		serviceClasspath := serviceJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "service"
		createSrcDir(serviceClasspath)

		renderPomFile("config/boot/static/service/pom.xml", &pomObject, servicePath)

		renderOtherFile(
			"config/boot/static/service/TestService.java",
			&pomObject,
			serviceClasspath+separator+"TestService.java")

		//===================================================================================================创建server模块 ===================================================================================================
		var serverModule = moduleNamePrefix + "-server"
		var serverPath = projectPath + string(filepath.Separator) + serverModule
		os.MkdirAll(serverPath, 0777)

		var serverSrcPath = serverPath + string(filepath.Separator) + "src"
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

		serverClasspath := serverJavaPath + string(filepath.Separator) + packageStr + string(filepath.Separator) + "server"
		createSrcDir(serverClasspath)

		//创建 config 配置java类
		createSrcDir(serverClasspath + string(filepath.Separator) + "config")
		renderOtherFile(
			"config/boot/static/server-boot/config/RedisConfig.java",
			&pomObject,
			serverClasspath+separator+"config"+separator+"RedisConfig.java")
		renderOtherFile(
			"config/boot/static/server-boot/config/CorsConfig.java",
			&pomObject,
			serverClasspath+separator+"config"+separator+"CorsConfig.java")
		renderOtherFile(
			"config/boot/static/server-boot/config/WebMvcConfig.java",
			&pomObject,
			serverClasspath+separator+"config"+separator+"WebMvcConfig.java")

		renderOtherFile(
			"config/boot/static/server-boot/config/SwaggerConfig.java",
			&pomObject,
			serverClasspath+separator+"config"+separator+"SwaggerConfig.java")

		//创建 controller
		createSrcDir(serverClasspath + string(filepath.Separator) + "controller")
		renderOtherFile(
			"config/boot/static/server-boot/controller/TestController.java",
			&pomObject,
			serverClasspath+separator+"controller"+separator+"TestController.java")

		renderPomFile("config/boot/static/server-boot/pom.xml", &pomObject, serverPath)

		renderOtherFile(
			"config/boot/static/server-boot/MainApplication.java",
			&pomObject,
			serverClasspath+separator+"MainApplication.java")

		renderOtherFile(
			"config/boot/static/server-boot/application.yml",
			&pomObject,
			serverResPath+separator+"application.yml")

		renderOtherFile(
			"config/boot/static/server-boot/application-prod.yml",
			&pomObject,
			serverResPath+separator+"application-prod.yml")

		renderOtherFile(
			"config/boot/static/server-boot/application-test.yml",
			&pomObject,
			serverResPath+separator+"application-test.yml")

		createSrcDir(serverResPath + separator + "logger")

		renderOtherFile("config/boot/static/server-boot/logback-prod.xml", &pomObject, serverResPath+separator+"logger"+separator+"logback-prod.xml")
		renderOtherFile("config/boot/static/server-boot/logback-test.xml", &pomObject, serverResPath+separator+"logger"+separator+"logback-test.xml")

		// 添加打包相关内容
		createSrcDir(serverMainPath + separator + "assembly")
		renderOtherFile("config/boot/static/server-boot/assembly/assembly.xml", &pomObject, serverMainPath+separator+"assembly"+separator+"assembly.xml")
		createSrcDir(serverMainPath + separator + "assembly" + separator + "bin")
		renderOtherFile("config/boot/static/server-boot/assembly/binTemp/start.sh.temp", &pomObject, serverMainPath+separator+"assembly"+separator+"bin"+separator+"start.sh")
		renderOtherFile("config/boot/static/server-boot/assembly/binTemp/stop.sh.temp", &pomObject, serverMainPath+separator+"assembly"+separator+"bin"+separator+"stop.sh")

		//===================================================================================================创建demo模块 ===================================================================================================

	}

}
