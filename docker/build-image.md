```
package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jhoonb/archivex"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func getDockerfile(path string) (string, string) {
	var dockerfilePath string
	var baoPath string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			name := info.Name()
			if strings.HasSuffix(name, "war") {
				dockerfilePath = path
			}
			if strings.HasSuffix(name, "Dockerfile") {
				baoPath = path

			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return dockerfilePath, baoPath

}

func main() {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	// 搜索打包的路径 获取dockerfile 绝对路径并打成tar 传递给docker引擎

	appPath, dockerfilePath := getDockerfile("/opt/bozz")
	fmt.Println(path.Base(appPath))
	appName := path.Base(appPath)
	fmt.Println(path.Dir(dockerfilePath))

	os.MkdirAll("/opt/bozz/", 0755)

	tar := new(archivex.TarFile)
	tar.Create("/opt/bozz/conf.tar")
	tar.AddAll(path.Dir(dockerfilePath), false)
	tar.Close()

	dockerBuildContext, err := os.Open("/opt/bozz/conf.tar")
	defer dockerBuildContext.Close()

	//buildArgs := make(map[string]*string)
	var arg2 string = "target/" + appName
	var arg3 string = appName
	args := map[string]*string{
		"source_war": &arg2,
		"target_war": &arg3,
	}

	options := types.ImageBuildOptions{
		//Dockerfile:     "./Dockerfile",
		Tags:           []string{"test"},
		Dockerfile:     "Dockerfile",
		NoCache:        true,
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		BuildArgs:      args}

	buildResponse, err := cli.ImageBuild(context.Background(), dockerBuildContext, options)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	//fmt.Printf("********* %s **********", buildResponse.OSType)
	response, err := ioutil.ReadAll(buildResponse.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println(string(response))
}
```
