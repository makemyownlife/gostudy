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
	"github.com/pelletier/go-toml"
)

func main() {
	fmt.Println("开始上传windows文件夹到相应目录")

	execDirAbsPath, _ := os.Getwd()
	log.Println("执行程序所在目录的绝对路径:", execDirAbsPath)

	config, err := toml.LoadFile(execDirAbsPath + "\\" + "config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		sftpClient, sshClient , err := connect(
			config.Get("sftp.remoteUser").(string),
			config.Get("sftp.remotePass").(string),
			config.Get("sftp.remoteIp").(string),
			22)
		if err != nil {
			log.Fatal(err)
		}

		var localFilePath = config.Get("sftp.localDir").(string)
		var remoteDir = config.Get("sftp.remoteDir").(string)

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
		defer  sshClient.Close()
		session, err := sshClient.NewSession()
		defer session.Close()
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		fmt.Println("开始执行后置脚本")
		myerr := session.Run("source /etc/profile && /usr/local/apache-tomcat-8.0.43/bin/startup.sh" )
		if myerr != nil {
			fmt.Println(myerr)
		}
		fmt.Println("完成执行后置脚本")

	}

	//睡一会
	sleepForAwhile(10)

	sleepForAwhile(2)
}

func connect(user, password, host string, port int) (*sftp.Client,*ssh.Client , error) {
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
		return nil, nil ,err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, nil ,err
	}

	return sftpClient, sshClient , nil
}

func sleepForAwhile(sleepTime int) {
	duration := time.Duration(sleepTime) * time.Second
	time.Sleep(duration)
}

func combinWindowsPathToRemoteDir(rootDir string, currentDir string) string {
	var remotePath string = strings.Replace(strings.Replace(currentDir, rootDir, "", -1), "\\", "/", -1)
	return remotePath
}

