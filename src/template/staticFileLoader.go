package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func readStaticFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

func createSrcDir(rootp string) {
	err := os.MkdirAll(rootp, 0777)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}

func copyStaticFileToTarget(filepath string, dest string) {
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	ioutil.WriteFile(dest, []byte(string(fd)), 0777)
}

func renderPomFile(pomSrc string, p *Inventory2, pomDest string) {
	var pomStr = readStaticFile(pomSrc)
	var pomTmp, err = template.New("pom").Parse(pomStr) //建立一个模板
	//将struct与模板合成，合成结果放到os.Stdout里
	var pomPath = pomDest + string(os.PathSeparator) + "pom.xml"
	var pomWriter, err3 = os.Create(pomPath) //创建文件
	err = pomTmp.Execute(pomWriter, p)
	if err != nil {
		panic(err)
	}
	if err3 != nil {
		panic(err)
	}
	defer pomWriter.Close()
}

func renderOtherFile(pomSrc string, p *Inventory2, pomDest string) {
	var pomStr = readStaticFile(pomSrc)
	var pomTmp, err = template.New(pomDest).Parse(pomStr) //建立一个模板
	//将struct与模板合成，合成结果放到os.Stdout里
	var pomWriter, err3 = os.Create(pomDest) //创建文件
	err = pomTmp.Execute(pomWriter, p)
	if err != nil {
		panic(err)
	}
	if err3 != nil {
		panic(err)
	}
	defer pomWriter.Close()
}
