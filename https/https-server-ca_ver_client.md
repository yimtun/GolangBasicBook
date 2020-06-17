
### server

```
package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "net/http"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
                   r *http.Request) {
    fmt.Fprintf(w,
        "Hi, This is an example of http service in golang!\n")
}

func main() {
    pool := x509.NewCertPool()
    caCertPath := "ca.crt"

    caCrt, err := ioutil.ReadFile(caCertPath)
    if err != nil {
        fmt.Println("ReadFile err:", err)
        return
    }
    pool.AppendCertsFromPEM(caCrt)

    s := &http.Server{
        Addr:    ":8081",
        Handler: &myhandler{},
        TLSConfig: &tls.Config{
            ClientCAs:  pool,
            ClientAuth: tls.RequireAndVerifyClientCert,
        },
    }

    err = s.ListenAndServeTLS("server.crt", "server.key")
    if err != nil {
        fmt.Println("ListenAndServeTLS err:", err)
    }
}
```


###  创建客户端私钥 证书请求文件  由 ca 签发



```
openssl genrsa -out client.key 2048
```


```
openssl req -new -key client.key -subj "/CN=tonybai_cn" -out client.csr
```


```
cat > client.ext << EOF
extendedKeyUsage=clientAuth
EOF
```


```
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile client.ext -out client.crt -days 5000
```





####  client

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

    cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
    if err != nil {
        fmt.Println("Loadx509keypair err:", err)
        return
    }

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{
            RootCAs:      pool,
            Certificates: []tls.Certificate{cliCrt},
        },
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



```
go run https_client_server_verifi.go
Hi, This is an example of http service in golang!
```


```
curl https://localhost:8081 --cert /opt/CloudLink/https-client/client.crt  --key  /opt/CloudLink/https-client/client.key  --cacert /opt/CloudLink/https-client/ca.crt 
```

```
Hi, This is an example of http service in golang!
```



```
curl https://127.0.0.1:8081 --cert /opt/CloudLink/https-client/client.crt  --key  /opt/CloudLink/https-client/client.key  --cacert /opt/CloudLink/https-client/ca.crt 
```

```
curl: (51) Unable to communicate securely with peer: requested domain name does not match the server's certificate.
```




```
curl https://192.168.20.133:8081 --cert /opt/CloudLink/https-client/client.crt  --key  /opt/CloudLink/https-client/client.key  --cacert /opt/CloudLink/https-client/ca.crt 
curl: (51) Unable to communicate securely with peer: requested domain name does not match the server's certificate.
```






####  重新生成服务端证书 以便使用ip 访问

生成服务端私钥和csr

```
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
```

```
echo subjectAltName = IP:192.168.20.133,IP:127.0.0.1 > extfile.cnf
```


```
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000 -extfile extfile.cnf
```


```
curl https://192.168.20.133:8081 --cert /opt/CloudLink/https-client/client.crt  --key  /opt/CloudLink/https-client/client.key  --cacert /opt/CloudLink/https-client/ca.crt 
Hi, This is an example of http service in golang!
```



####  客户端修改


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

	cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://192.168.20.133:8081")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```
