```
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	httppost()
}

func httppost() {
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)
	cliCrt, err := tls.X509KeyPair(certPem, keyPem)

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

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	data := `{"user_name":"yandun","pwd":"123456"}`
	req, err := http.NewRequest("POST", "https://192.168.3.3:8081/auth", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//post数据并接收http响应

	respBody, err := ioutil.ReadAll(resp.Body)
	if err!=nil{

	}else {
		fmt.Printf("response data:%v\n", string(respBody))
	}

}

var caData = []byte(`-----BEGIN CERTIFICATE-----
MIIC/zCCAeegAwIBAgIJAKL+fvYeq6a9MA0GCSqGSIb3DQEBCwUAMBYxFDASBgNV
BAMMC3RvbnliYWkuY29tMB4XDTIwMDUyOTEzMDMyNFoXDTM0MDIwNTEzMDMyNFow
FjEUMBIGA1UEAwwLdG9ueWJhaS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQCuIgtaFsE9hlc8/aqnRQ2c3xH9SphTKno+dSgoN0GbXakvVRCToRbc
E15HtqlbtXqgnwpffC5OqePpG8S0wH6TufvA0k5MQ4W7/YLwJfoQgKkafzsYSz6U
5deTn0+CxTeVvEUfi1TaGDWjf2gNhlxu35ESPC6kjErb1wdAHiJVDWWpSIRFAoYD
mmjFXchjnsqHHzRG5J1/7XP9xuH8CPktt1Oiq7YBNUyl45zyEBWryCddZV+kycP9
GkzSUHZNXNPHv6uNwlr5Cm4tLh6B/yAfSCGFcNVtPhcX28mqpgdyojr6WMIoHkDb
FixTBTdGwJS339kuAouFgxI44RhkIcSfAgMBAAGjUDBOMB0GA1UdDgQWBBSbfKHS
75POY4Myfxu+6y/swMEh5TAfBgNVHSMEGDAWgBSbfKHS75POY4Myfxu+6y/swMEh
5TAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQAztRJ6G4eHwCDRDorg
b6RhKPsgqBAmZuWESOt77MPtKlwBdxBclcmuIfmMQs7OBxN/qhkB1yql9yvcsOR4
pn4dmwPRyi4T0sZBOhiIcIMd6E7Zsn620Gj+vsBRqpvTeu4W3briX+X70xICQnx4
uEclS4pCo3wJGXz0mz7X27b6TDdaFsxTPi0QZxkno66NAQLO+PCplDTYFZVeDWeP
yjlkuV0qc5Jjf71Tndq6ifPfYvkmchdL3Uj/0NwGJiSXE9DAmd0lzwP0bM9thJg+
dFBXtlbq/p1q2TIRJgVXKnMLjyu8e6KDO9HgojdJLmylP+3XKxewa5UdW2Hd2wqY
DqLb
-----END CERTIFICATE-----`)

var certPem = []byte(`-----BEGIN CERTIFICATE-----
MIICxTCCAa2gAwIBAgIJALEGoUYJ94AXMA0GCSqGSIb3DQEBCwUAMBYxFDASBgNV
BAMMC3RvbnliYWkuY29tMB4XDTIwMDUyOTEzMTQyNVoXDTM0MDIwNTEzMTQyNVow
FTETMBEGA1UEAwwKdG9ueWJhaV9jbjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBANjq5z2d6iBeMsaP/AAZQM5LKiBwBkUu1dy2ej6C4NLKctUAT1QqgvTz
PbfQMfcUUmhNcDxz142dqzwnr4eFxWO5S9rHUwulr9CqBDRIi2CByIjJnhVgSNq1
7CWmTFTwygu82L4jH/Vjw2ZFcxMYRetCoU6Oc2nBnMlI9lr+DcHmkTsWCH3tIx31
Zbb6Hz5U8xmEjhM5JB56xKLNKt6vD7P/qPNvAwzaLpEf/yyPcTEUaa1S8ZlAgJgu
2RnrzTchDyNVfe9q5hx3STbmdfAv6qLu0hdw4GRRezfzi4gxu+Tmu9xPt/Oxz/XJ
pv68MAfDE+uhnNOUSADkfmgAr392zzMCAwEAAaMXMBUwEwYDVR0lBAwwCgYIKwYB
BQUHAwIwDQYJKoZIhvcNAQELBQADggEBAApVUZAC6/+PKqStMx7gWObZql2yW8oa
fF9Pv74X2HUz/FFKl98vhcrPtRguni3qCdA8WkGU0SFi6QHLKBX+Qxa6eF3OU79e
G+w+aDR+FSl8eiLc+CjzStlHcU2TdlPNSF+kLVnjYeuU54CXuE6NHgg0G3zQsdQE
DKPU2RR32eGNajGvsljcylavYwKvLKKq70f+HeQO95ZB3yGj18tWojX5fQO9J3kC
SYne3A0AjYX6fYMF0RB2s4VrZNd/7XeSUr4tEDXmmyaNeNapOxaM614s+9N6T0oB
lj0AHA8ORyzUE0bdRWJEgUKg+n7lrAFngwpQ9DvhN2DNscZtw4weUWQ=
-----END CERTIFICATE-----`)

var keyPem = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA2OrnPZ3qIF4yxo/8ABlAzksqIHAGRS7V3LZ6PoLg0spy1QBP
VCqC9PM9t9Ax9xRSaE1wPHPXjZ2rPCevh4XFY7lL2sdTC6Wv0KoENEiLYIHIiMme
FWBI2rXsJaZMVPDKC7zYviMf9WPDZkVzExhF60KhTo5zacGcyUj2Wv4NweaROxYI
fe0jHfVltvofPlTzGYSOEzkkHnrEos0q3q8Ps/+o828DDNoukR//LI9xMRRprVLx
mUCAmC7ZGevNNyEPI1V972rmHHdJNuZ18C/qou7SF3DgZFF7N/OLiDG75Oa73E+3
87HP9cmm/rwwB8MT66Gc05RIAOR+aACvf3bPMwIDAQABAoIBAE2hE4yLnSRkUcuZ
79ehMf4iDDt7m+LadhL9AxaoBAmsBhiQedNnO2KqQmYsDhWcojlNrWMVGglGaeen
VDwt9YxQ3gK33tKGQztr/QExDSBGkhvR+88unCQZgx+eOicT3VsCBJdXPrK1N1K8
fAhTYYJmYBAbDtbEQXYPtIqv7QB8+JQzoYj4r6uC/6bYLFc1BL3AizSxHvNQ6qyS
jVLNCw/b2XSV/Gct+JO25lMR97eTtUZDO+aCJhRHtf2yH4AVJ2kjTSDX0zAwvEyn
kTqVOhoBQ9UuHuDxBDY4evXSwm/HPtnLKt+n3VJm7hNiMbkfeo5GGRPBgsgS3SXn
PBjgqjECgYEA9y6corfZu/AqxxnvDi1wBreK1MxRoSI6sOPqusPX/i6bScnG4khH
hfIJsQtjjz6yYoB/N/Ph6XKPh6FwaICmZKVmiQafDy1wdGhBnjKZUX2+z6rGW82C
azEhXaCvsQrP+8dlDQwLKX9W6wl+msfPigtOgHoW+UjTKx24VLhZl00CgYEA4Kfm
uZKIZ5drXN0BxqeHKgtAueWEbeonmxjVGdcCqhSyIzZ/tb6hjKOL0MKqkzug9Ful
kymdaR/tUFHzbF91MTOSooWAz0VuNm/9msooSLA6wGfixeoijLS2AXoe/zuE/BpO
+BCYjyrN+m6RjQFqzJdiM8vqbfvrSMJ+/rg3wH8CgYEA9LRAfRZlh5kLwG/8KZzx
+rRZzLE5Xrv0PopkJuEbQ5gAQ9xt+DZ2GDS1Q4By36zCDKIJT+sxC1tOnIdGMS6r
i5fzGlRjwzLI0XUlAOg2wGYsMIpwNRIOuu+4iGg7hpmdrTSI5ZG2uy5FbgnIPJLy
80aeQLKda8ycwWJ6yeR3TUECgYAJWtR/u7tUXPjKiCVvyUAWnxtV4lnOXebCTZlq
Fwr0XCJC1x0gS0zc116WFdUWhnAgXMSbSStV2k9+fbseMydlIhmkDN0dRG3fdHBi
BrMAu93mby0BQsungkrl6tjplefmu424RLClnYmwOPR0o693z4H2WDJt3ASjJf41
yyP/LwKBgQDiS6GRXSX59WkMxQREnpInH25RKLdowBh1lDI9ZsRVWEWGGkwJwhCA
QMcNGD/Z6XTnk3w22u66N69dcefteAOC4425inNRcG5MM3s3F+xJ3i/RHu157zIj
3RfPQhGacZLLy4EKBeL2FUH75T2Qzn3H32h3MaALPc6DCj9HjLRsqQ==
-----END RSA PRIVATE KEY-----`)

```
