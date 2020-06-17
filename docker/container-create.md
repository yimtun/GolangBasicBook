```
package main

import (
	//"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	//
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	//create container
	ctx := context.Background()
	nport, err := nat.NewPort("tcp", "9001")
	if err != nil {
		fmt.Println(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "172.16.100.216/yygl/lz-eoms-appserver:v1.1.0",
		Tty:   false,
		ExposedPorts: nat.PortSet{ //docker容器对外开放的端口
			nport: struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			nport: []nat.PortBinding{nat.PortBinding{ //docker容器映射到宿主机的端口
				HostIP:   "0.0.0.0",
				HostPort: "9001",
			}},
		},
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/opt/config/appserver/application.properties",
				Target: "/opt/conf/application-env.properties",
			},
		},
	}, nil, "appserver")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

}
```
