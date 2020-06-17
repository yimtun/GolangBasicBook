
```
openssl genrsa -out ca.key 2048
```

```
openssl req -x509 -new -nodes -key ca.key -subj "/CN=tonybai.com" -days 5000 -out ca.crt
```

```
openssl genrsa -out server.key 2048
```

```
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
```

```
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
```



#### server 

```
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,
        "Hi, This is an example of http service in golang!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":8081",
        "server.crt", "server.key", nil)
}
```




#### client


```
package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    pool := x509.NewCertPool()
    caCertPath := "ca.crt"

    caCrt, err := ioutil.ReadFile(caCertPath)
    if err != nil {
        fmt.Println("ReadFile err:", err)
        return
    }
    pool.AppendCertsFromPEM(caCrt)

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{RootCAs: pool},
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get("https://localhost:8081")
    if err != nil {
        fmt.Println("Get error:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```
