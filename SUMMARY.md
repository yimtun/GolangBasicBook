# Summary

* [二进制](jinzhi.md)

* [GO基础](README.md)

* [IO](IO/IO.md)
    * [[]byte2io.Reader](IO/byte-Slice2-io.Reader.md)
    
* [OS](OS/OS.md)
    * [exec](OS/os-exec.md)
    
* [文件查找](file/find.md)

* [数据基础](basic/basic.md)

* [docker sdk](docker/docker.md)
    * [go mod](docker/go-mod.md)
    * [connect localhost](docker/connect-localhost.md)
    * [tls hard code](docker/tls-hard-code.md)
    * [build image use args](docker/build-image.md)
    * [container create](docker/container-create.md)
    * [pull image](docker/pull-image.md)
    * [mount args](docker/mount.md)
    
* [k8s sdk](k8s/k8s.md)
    * [go mod](k8s/go-mod.md)
    * [kubeconfig 硬编码 yaml 格式注意 set paste](k8s/client_kubeconfig_hard_code.md)
    
* [ssh  client](ssh-client/ssh-client.md)
    * [输出到标准输出](ssh-client/ssh-client-stdout.md)
    * [输出到字符串](ssh-client/ssh-client-stdout-str.md)
    
* [https 和证书](https/https.md)
    * [一个简单的https 服务器](https/https-server-simple.md)
        * [https 客户端](https/https-clinet-simple.md)
            * [https 客户端 忽略验证服务端](https/https-clinet-simple-no-ver.md)
            * [使用服务端的证书验证服务端](https/https-clinet-simple-ver.md)
        * [服务端使用ca签发证书](https/https-server-ca.md)
            * [服务端验证客户端证书](https/https-server-ca_ver_client.md)
        * [客户端证书hard code](https/https-client-hard-code-tls.md)
        * [https client post json](https/https-client-post.md)
    * [https 双向认证 路由](https/https-server-router.md)
    
* [test](teshu.md)

* [string byte](byte.md)

* [json](json/json.md)
    * [将数据转换为json](json/data2json.md)
        * [struct 2 json](json/struct2json.md)
        * [map 2 json](json/map2json.md)
        * [(string key string vlaue map) 2 json](json/string-map2json.md)
        * [slice 2 json](json/slice2json.md)
        * [序列化slice 元素是struct](json/slice-struct2json.md)
    * [反序列化json](unjson/unjson.md)
        * [json2array](unjson/json2array.md)
    
* [etcd](etcd/etcd.md)
    * [put](etcd/put.md)
    * [get 判断一个key是否存在](etcd/get.md)
    
* [vim-go](vim/vim.md)
    * [vim8 install](vim/vim8-install.md)
        * [vim-plug vim-go](vim/vim-plug.md)
    * [go install](go-install/go-install.md)

* [常用库](libary.md)   
    * [**pflag**](libary/pflag.md)

