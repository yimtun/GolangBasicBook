```
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"os"
	//tls
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

func yyglgetTLSConfig() (*tls.Config, error) {
	var _tlsConfig *tls.Config
	if _tlsConfig != nil {
		return _tlsConfig, nil
	}

	cert, err := tls.X509KeyPair(yyglcertPem, yyglkeyPem)

	if err != nil {

		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(yyglcaData)

	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	return _tlsConfig, nil
}

func main() {
	ctx := context.Background()

	tlsConfig, err := yyglgetTLSConfig()
	if err != nil {
		fmt.Println(err)
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	httpclient := &http.Client{Transport: tr}

	//改成自己的ip 加端口
	cli, err := client.NewClient("https://172.16.99.73:2376", "1.13.1", httpclient, nil)
	if err != nil {
		fmt.Println(err)
	}
	// pull
	authConfig := types.AuthConfig{
		Username: "yygl",
		Password: "Yygl@123",
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	//reader, err := cli.ImagePull(ctx, iname, types.ImagePullOptions{RegistryAuth: authStr})
	reader, err := cli.ImagePull(ctx, "172.16.100.216/yygl/lz-eoms-appserver:v1.1.0", types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)
}

var yyglcertPem = []byte(`-----BEGIN CERTIFICATE-----
MIID5zCCAs+gAwIBAgIJAL4gbGex+AAXMA0GCSqGSIb3DQEBCwUAMBkxFzAVBgNV
BAMMDjE5Mi4xNjguMTAuMjUwMB4XDTIwMDYwMzE0NTQ0MFoXDTIxMDYwMzE0NTQ0
MFowETEPMA0GA1UEAwwGY2xpZW50MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIIC
CgKCAgEA2lWlYuSLv8KtIEe1PYw8Yu474m65JhEVF5019on1X7l0AsOOQtdpsa4j
UxzXIs4rl+RAGlhP2i8+rM/oMGAsgiglxyRbpZb2ehJzp//LJ5MBHyJx4kWbuUnr
nolXZiPMnMK1f62Inr+f7+JvFw4Tgk6+QuAYmCVKDrPvIT6+SEmI7VrDry91WX3W
RvCbUoD/TNe+LaTzLU/WVKgNKe/GDAtflqBcKspUecCBs++Fxs1N3USIzOFgHOk9
Wai8J7epDA2+f0f8fqxUmTk+Z/tbCDiPcSqpJCTyDrlmOTURCP43lxqg01qAJtqs
98uej+y7RTTe1qunr8fJ6QCccVpfzEzFgfuYACrTeqWBDcFNZO+bppQjHoN0WAPp
fgMVGyk5v72Jp9d85uGNsY5nz54U3H3WLuPVjsaDzvCQDAkDsSvIpVqHV2XL95Id
5PNooXsLjFWzK8Xtnh9Ba/B8P5QKr4i+sZfzhZGCDy2bvl3ufK/OQR4LFPjq8jNs
3ftCHHBhknKf5uN7o82MaO8SQeoRkIuuom6tenZdQNAsd+gMFx+okc01ZR9ta6Um
Pgz+y4+81Yo9EoN8AnfGY2S727DHlfNQ9A+kBoajNM+yHOuR/lQH7NyRQsqOvXxP
KAoM4DsLOSxBtmVCMF+TbZiZznxDaOla18mL2F7F3x1bkKB5g9UCAwEAAaM6MDgw
IQYDVR0RBBowGIcErBBjCocEfwAAAYcErBBjSYcErBBjSjATBgNVHSUEDDAKBggr
BgEFBQcDAjANBgkqhkiG9w0BAQsFAAOCAQEAluKDXfHwVBOIajRd1ClmBfSNDeog
+65j0GVLPnEXh9tdYJ9BDGNH3tOO1+W5ehb31XjbQGDGsUz9NcPDBtisF/16zWzG
09P+bdcq56udEr8QOY2RXYQJodr6AC8t4RmftWAFCfE7QqYJz2H0Ame7yvw22LIU
kFp7c4Q6i3AVSCN0gC/CUyW5cwjvuMfed5ZdNIIgRJ/O8z/w8baBa/HhSUEas94N
JEeAfl71VO+UhbhbxU3TFhMfYwiZv2+FKnYQI/TIojrgYgVzAET+vQfyaFoYsloi
2oSMc+EV7cxX2TfMG7W3S1QYxNrmeWodPYBYX1GVgFI/0UqqFjQR8p3EOg==
-----END CERTIFICATE-----`)

var yyglkeyPem = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

var yyglcaData = []byte(`-----BEGIN CERTIFICATE-----
MIIDBTCCAe2gAwIBAgIJAPs7Pe27NKzUMA0GCSqGSIb3DQEBCwUAMBkxFzAVBgNV
BAMMDjE5Mi4xNjguMTAuMjUwMB4XDTIwMDYwMzE0NTMwNloXDTQ3MTAyMDE0NTMw
NlowGTEXMBUGA1UEAwwOMTkyLjE2OC4xMC4yNTAwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDocld1e168wxhUlcXCHmpctqgMjhymkSW5rUMcAG6fGvYj
3U7ohwfV+/kaKoCsGg1nyUcpa26pIMEXvRgx7aBEo28x27uqV5r3pipAmfNbuJvb
ZktrmKyxo1Z6r1euZ3JFB9vnuGOVYXbtJhpMbpi+7H8ZIDcjgmcR1ijF1l5zDjdP
OJDbed0L48OARsJuky6zXKtJCLRVGop5TNtswcmze1xbiy01z1eIirtIZaoZBg8G
/c0KEEhMw1pjzLVKyiuzlhCYxE4Qs3kI0GlnUW3Qe+QTjqLYCskAzVS94RFnbwkV
lyHq+/9LCmnatUynIGpnlMysCJ+kHkMXlqWH1dZXAgMBAAGjUDBOMB0GA1UdDgQW
BBRpPnLo3/QzOa1dTqZnmiJMWATCHzAfBgNVHSMEGDAWgBRpPnLo3/QzOa1dTqZn
miJMWATCHzAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQDW69UDptu6
KD5YbPpjhdwl6eUodwWEFoWS1/UiYwu0tlhClUsiozyqzd01lTWW0rLftTp3K5Tk
MqpIJrcrK/hZlbqdXGPyKywJvpuEzmnyUaU1diOA8Muac6Hn6GdNyOvuS2roisv3
lQyKmKldDvhSqCoDvRwss9BMjGFqUK3H4e5De0NNoBLt8Ww+LFoFUz0C4Xwp4n91
yuecXXdelK101JccY5YVG8RgxT/jbE349XaC89G448JSTwbqBHIDWxgbKlkdHAA4
5eC7eE2slj4WDbDGScldjycET6j+ILF6yFlwnDl3padTvQs/eYlKX1H2W7PBs9VK
IlwPO9WsPlC8
-----END CERTIFICATE-----`)
```
