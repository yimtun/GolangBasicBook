创建私钥

```
openssl genrsa -out server.key 2048
```

创建证书

```
openssl req -x509 -new -nodes -key server.key -subj "/CN=yandun.com" -days 10000 -out server.crt
```


```
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,
        "Hi, This is an example of https service in golang!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":8081", "server.crt",
                           "server.key", nil)
}
```


```
curl  https://192.168.20.133:8081
```

服务的日志

```
2020/05/29 20:09:09 http: TLS handshake error from 192.168.20.133:36738: remote error: tls: unknown certificate authority
```







