```
package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	//
	"crypto/tls"
	"crypto/x509"
	//	"io/ioutil"
	"net/http"
)

var certPem = []byte(`-----BEGIN CERTIFICATE-----
MIIGyjCCBLKgAwIBAgIJAL03tsQihQYCMA0GCSqGSIb3DQEBCwUAMEgxCzAJBgNV
BAYTAkNOMREwDwYDVQQIDAhTaGFuZ2hhaTESMBAGA1UECgwJV2hvQXJlWW91MRIw
EAYDVQQDDAlsei1kb2NrZXIwHhcNMTkxMjA0MDMyOTU2WhcNMjAxMjAzMDMyOTU2
WjARMQ8wDQYDVQQDDAZjbGllbnQwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK
AoICAQC9b+fk6DCwhyODdzbg7mNXBScW6aWQDPxtoRKAIqC12+G9VT3uI5/EHqqW
OL5tu5sYkxGJzXqhvGepctiKXV0I0aesRXbVErhNY2VQiw0wzzk7hc36PqDYRwTR
TcYhE86757hVsKG+qa9BIWJeTqKPFHPSLAvwymNAZRJIVZvkdWXc18LsqOKlZN7n
wwPuBCyniVZiGH7I2kYudX1oEy61VmqVj+vCPJuVqPceb44oabvf3fURQalIiukb
lFTmRGfSfiLjXMZPFji/d2t5++wo0ZmtzfuyAJsP6NnMK0yWmZ1Rcyaud0f1O3cs
AD4uohUT2PBeufc3ciV663U5aevjXVBQcwjvVBLJMbzf7vDRdRtjLu2Hd7SQiHur
4GIpK3TxPuh938aFMhFbob+sxrjzd3ITYx4HXEUZworDslE/idNCs2JxJfKJoYdh
LZHJ0JFeEwL0E8KAMh5g2KNbgxDhhRnPqlwWXMoT3zlUWM9a1VZh2+FDbyquiHCC
VX4sQ2nOy5uluZjBe863b/mxSV13FRdDSyZCfsPL6BD1970tHyyd0zhvi77+LwkJ
Kuxnf18VQzCceVyveCyV3E8f2iD+kOae6h8/i9TYPhEmkagTkQLRR6zWY7QP6bio
B8SeXL6Egn7lhcENyLza9ZCqp7MiPQckLsash3cF6y9E9QWr+QIDAQABo4IB7DCC
AegwEwYDVR0lBAwwCgYIKwYBBQUHAwIwggHPBgNVHREEggHGMIIBwocErBBk2IcE
ChQAyYcEChQAyocEChQAy4cEChQAzIcEChQAzYcEChQAzocEChQAz4cECv+A9ocE
Cv+A7YcECv+A44cECv+Az4cECv+A94cECv+A4IcECv+A6IcECv+A7ocECv+Ak4cE
Cv+Aa4cECv+AzocECv+AuYcECv+AyIcECv+AvYcECv+A14cECv+A1YcECv+ArYcE
Cv+AfIcECv+Au4cECv+AnocECv+A5ocECv+AhocECv+AeIcECv+Ad4cECv+AoIcE
Cv+AuocECv+AzYcECv+AsIcECv+ApocECv+Ag4cECv+ArIcECv+AlIcECv+AkYcE
Cv+Al4cECv+AhYcECv+AkIcECv+AgocECv+AyocECv+AmocECv+AeYcECv+AeocE
Cv+AfocECv+AgIcECv+Ae4cECv+Af4cECv+AfYcECv+A3YcECv+AlocECv+AhIcE
Cv+AjYcECv+AiIcECv+An4cECv+AdIcECv+AdocECv+Ab4cECv+AjocECv+AcocE
Cv+ArocECv+A4ocECv+AcIcECv+AvocECv+AdYcECv+AcYcECv+AbIcECv+A1ocE
Cv+Ac4cECv+AhzANBgkqhkiG9w0BAQsFAAOCAgEAN6nRSmOp3Rayr2CZAWPm8Km0
rUCpKGKxuDO8cfZ7+NOPCJioBrazzJViEIGJqOeUkxy06od6RqwBkd/ja8u0fvAy
Sq4+ycTJjlQTLie29e1a9DK1doXr+UjmlGf9BMWsQxoVy8DXS3odXyrrv/KcH3Hu
NXSc55vNzN/XydLlrTo0RsDX42W9X6+3BeBNALM096reNU6aXkon8xYS9ZGv5bYw
vmOI0bhpvz8Uz1yGnTvYXNtSaTUZ0klHyCJh7/WgEC12o5LE9BdWwRvZU3/Emm2D
C2Ykg1zPZFGfQRD1NQbzxKuZfqpaWW0pwcmOxa6FZPjlsYcjSeLSf4woiGrt65sc
fyqNrJ0e4vxRFCTg4puI3Vy4X7Pa0GuM8s+Az7fs54hUoPN2B5KyLbhvUmhh3C2+
7/fVenGzIMFQ8FVHwCix96jqYHq5Frl0B59fhCsR6aBTnzMMRuhM83dlaIYWHpCX
AQYyx2yuZuI7+INqHxJ1A24+kXEXgwWLs+o2D/Xf7g/OJ8u8UQBvKHOz7VLUSoMG
j8a/CVr5b5RnEX2eJ9PqHpPsrWbdbbhkLF193UVv1nezTYUSZDU/a8R4KCo1Pqlr
gPVVtDA/ni6VwCAizT4JSWlbleca9kZ4jTfW9+H2l1UPOIKXQ4m4jQnUGmCVnU4F
ty5Z9eCMSkHmG+aVedY=
-----END CERTIFICATE-----`)

var keyPem = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

var caData = []byte(`-----BEGIN CERTIFICATE-----
MIIFYzCCA0ugAwIBAgIJAMnXOfSPRBhjMA0GCSqGSIb3DQEBCwUAMEgxCzAJBgNV
BAYTAkNOMREwDwYDVQQIDAhTaGFuZ2hhaTESMBAGA1UECgwJV2hvQXJlWW91MRIw
EAYDVQQDDAlsei1kb2NrZXIwHhcNMTkxMjA0MDMyNzMzWhcNMjAxMjAzMDMyNzMz
WjBIMQswCQYDVQQGEwJDTjERMA8GA1UECAwIU2hhbmdoYWkxEjAQBgNVBAoMCVdo
b0FyZVlvdTESMBAGA1UEAwwJbHotZG9ja2VyMIICIjANBgkqhkiG9w0BAQEFAAOC
Ag8AMIICCgKCAgEAr9rGOUBh8mCqGCvBg2q6mhSMIsXi12h4aHZHsrFrzCnnF9Wg
nqX6tw9cThu5v3dWN9l4xgGQOkb17UnIKkuWMgSo/pa2OTGwBe9R3VqOVzCgft8o
qOkgV/aFswSIpweJh6ftcLvq7PJNpkevbZm3aWFzU3306nx+RKMSfHmV/beoIST9
/b8xnWoLePdsmrYG3QSMTvFCP0pJSHleCl0JuncXD4PIEiyY+sikbfP0kdIu4IDy
Cz54/AvherKOnhrzw2fN49SAmMSStFdIRQl4IKtaiywzKct9xFGCsOvDizu5trD3
CLRUb4HDoQw56nQlTwOzS7o1DQ33heyenZ1D6mzriz3r2mV4IMRqZQ8xDEVCTZkk
5ihYBkaKPKiJwuxTwxQ/AF//s2pur2Hq5UEdk7irBTelvmmFnvVazWCWiRh4SHHY
hMLZG3MR8uPOQv0FAAaLecN97F78X9uTbfyxpTVTBDCgAMWUdFHXpdo0rVKSxm1N
2IGecB2bCcYPKo46gAj7UHAJic+OWjNXtFs4fqE4CMCIGijOyUbxaiQfliZPn41r
OB8R5EXG2xAhShAVnEVgKoPSFtXFeDjLlNGBob2QNS3epqEC3YsktHmmX+UWyZNr
gVfC+ptiLRwwsREThmlgmvxrG2Zwbp2gp2Sjx4rogQYVIzSSItY8Oa5wymUCAwEA
AaNQME4wHQYDVR0OBBYEFEQJ9ivTuxyIPO5o5leSuD0bWtw0MB8GA1UdIwQYMBaA
FEQJ9ivTuxyIPO5o5leSuD0bWtw0MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEL
BQADggIBAHMoMwh2EUoX0vQBWVBTtWBWBO/gz2Ht+Q6j1qjQA7qGo5fu9msmMvgm
XmHq/mdhAtYUCK38y9H/Z5Xg3eoKA/fW2VhehB4v/D1npira4NW7BEK3Lfr8Tu2j
ZSEpI82dk673gYRv2kJa9qoIO31TFMQ8qpySkOrJz/IxXd4vwFQx42jAusyGdHzb
6XseqDKhMIt09mx+PjcDSijAIi9AauLqsTbrK3Xbgbb68jdfqjbkRedc1/RuLeLj
K4NpzDTZ/fq/NBt/S8IjqY/99vjUbEadkN2XlZGjIP894uVJmsRTea5Xtfbim+vY
8wBYvTwPUxcZSXgOlU9UgH6JClJ3t/htxQt7p5GBG7tu153JeboNuJ3oPwWnwHkw
puWv5P0zAAHoOksEL3EB1NlEVW0QcPg1ba9YA1tktFutlVO9vwDKlW3pLr0TJSy3
tlV8pSK8ksAlPoI30TWCRFFNfOZqvTorR390xDooKREC2Wr0YFgrI5opclUgBCOq
zX3io6J9JXTCRFFyrCXlRYECQW3yr2ez67qfmjkFQnz0Jf8ow4mU94w2rZNpGdNY
KJlAZLDfuEis5OHPV0MTc96QhviF0YFURxRQ6mnPlGwcqIOo8KuMe5YfQNkLuHQ+
oEwmxRfh4AwuNnSueoSjummSmCVC8ICizkDzNBL74r4gfOeKdSwg
-----END CERTIFICATE-----`)

var _tlsConfig *tls.Config

func getTLSConfig() (*tls.Config, error) {
	if _tlsConfig != nil {
		return _tlsConfig, nil
	}

	// load cert
	//	cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	cert, err := tls.X509KeyPair(certPem, keyPem)

	if err != nil {

		return nil, err
	}

	// load root ca

	//caData, err := ioutil.ReadFile(wechatCAPath)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	return _tlsConfig, nil
}

func main() {
	tlsConfig, err := getTLSConfig()
	if err != nil {
		fmt.Println(err)
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	httpclient := &http.Client{Transport: tr}

	//改成自己的ip 加端口
	cli, err := client.NewClient("https://10.255.128.147:2376", "1.13.1", httpclient, nil)
	if err != nil {
		fmt.Println(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}
}
```
