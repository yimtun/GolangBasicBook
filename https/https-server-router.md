```
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
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

func aServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "a!\n")
}

func bServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "b!\n")
}

func main() {
	// 注册路由
	http.HandleFunc("/a", aServer)
	http.HandleFunc("/b", bServer)

	//
	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr: ":8081",
		//	Handler: &myhandler{},
		Handler: nil,
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



```
curl https://192.168.20.133:8081/a  --cert /opt/CloudLink/https-client/client.crt  --key  /opt/CloudLink/https-client/client.key  --cacert /opt/CloudLink/https-client/ca.crt 
a!
```

```
curl https://192.168.20.133:8081/b  --cert /opt/CloudLink/https-client/client.crt  --key  /opt/CloudLink/https-client/client.key  --cacert /opt/CloudLink/https-client/ca.crt 
b!
```
