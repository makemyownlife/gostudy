package template

import (
	"fmt"
	"io/ioutil"
	"os"
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
