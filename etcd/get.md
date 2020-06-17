```
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.16.99.102:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	defer cli.Close()
	//设置1秒超时，访问etcd有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), (time.Duration(10) * time.Second))
	//操作etcd
	//取值，设置超时为1秒
	//resp, err := cli.Get(ctx, "/xxx")
	resp, err := cli.Get(ctx, "/logagent/conf/")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	fmt.Println(len(resp.Kvs))
	/*
		for _, ev := range resp.Kvs {
			fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		}
	*/
}
```
