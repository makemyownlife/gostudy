package main

import (
	"fmt"
	"time"
	"log"
	"golang.org/x/crypto/ssh"
	"github.com/pkg/sftp"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("开始上传windows文件夹到相应目录")

	sftpClient, err := connect("root", "123456a?", "192.168.1.205", 22)
	if err != nil {
		log.Fatal(err)
	}

	var localFilePath = "D:\\server\\lib"
	var remoteDir = "/opt/lib"

	//先创建远程相关的目录
	filepath.Walk(localFilePath,
		func(currentPath string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			var transferFile = remoteDir + combinWindowsPathToRemoteDir(localFilePath, currentPath)
			fmt.Println("开始创建远程文件:" + transferFile + " 本地文件：" + currentPath)
			if f.IsDir() {
				sftpClient.Mkdir(transferFile)
				return nil
			} else {
				srcFile, err1 := os.Open(currentPath)
				if err1 != nil {
					log.Fatal(err1)
				}
				defer srcFile.Close()
				dstFile, err2 := sftpClient.Create(transferFile)
				if err2 != nil {
				}
				defer dstFile.Close()
				buf := make([]byte, 1024)
				for {
					n, _ := srcFile.Read(buf)
					if n == 0 {
						break
					}
					dstFile.Write(buf)
				}
			}
			return nil
		})

	defer sftpClient.Close()

	//睡一会
	sleepForAwhile(10)

	sleepForAwhile(2)
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func sleepForAwhile(sleepTime int) {
	duration := time.Duration(sleepTime) * time.Second
	time.Sleep(duration)
}

func combinWindowsPathToRemoteDir(rootDir string, currentDir string) string {
	var remotePath string = strings.Replace(strings.Replace(currentDir, rootDir, "", -1), "\\", "/", -1)
	return remotePath
}
