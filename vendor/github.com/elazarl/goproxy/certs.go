package goproxy

import (
	"crypto/tls"
	"crypto/x509"
)

func init() {
	if goproxyCaErr != nil {
		panic("Error parsing builtin CA " + goproxyCaErr.Error())
	}
	var err error
	if GoproxyCa.Leaf, err = x509.ParseCertificate(GoproxyCa.Certificate[0]); err != nil {
		panic("Error parsing builtin CA " + err.Error())
	}
}

var tlsClientSkipVerify = &tls.Config{InsecureSkipVerify: true}

var defaultTlsConfig = &tls.Config{
	InsecureSkipVerify: true,
}

var CA_CERT = []byte(`-----BEGIN CERTIFICATE-----
MIIF9DCCA9ygAwIBAgIJANbTQLWhIpwuMA0GCSqGSIb3DQEBCwUAMIGOMQswCQYD
VQQGEwJJTDEPMA0GA1UECAwGQ2VudGVyMQwwCgYDVQQHDANMb2QxEDAOBgNVBAoM
B0dvUHJveHkxEDAOBgNVBAsMB0dvUHJveHkxGjAYBgNVBAMMEWdvcHJveHkuZ2l0
aHViLmlvMSAwHgYJKoZIhvcNAQkBFhFlbGF6YXJsQGdtYWlsLmNvbTAeFw0xODA2
MDExMjU4MjJaFw0zODA1MjcxMjU4MjJaMIGOMQswCQYDVQQGEwJJTDEPMA0GA1UE
CAwGQ2VudGVyMQwwCgYDVQQHDANMb2QxEDAOBgNVBAoMB0dvUHJveHkxEDAOBgNV
BAsMB0dvUHJveHkxGjAYBgNVBAMMEWdvcHJveHkuZ2l0aHViLmlvMSAwHgYJKoZI
hvcNAQkBFhFlbGF6YXJsQGdtYWlsLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIP
ADCCAgoCggIBANGltmmM9u4qNXDq1PevbmMMyjiy3SejaNDLQNydx0nhRqjV+G/P
ruLfvLuj97BqCq4WnKl3mA5t3hcJYTrSdfKbG20giaM8aoN5xLUBs1PE4pp4Agaq
kpSjNpnep+Lq7zBsBNZZBIdZsPNzRd6P4UjzQWKekFGymtWuO10niCxt/R++3but
Uc0Pc3t4wbA2/YI5ShUygMWgZ8TdxkY7dUj3OnlF02QLSl2v7lvhoznPVJ5gUS/n
Vw6fgzm29NEZryI5lukjUI68LOQ9ZKeJ++ObHaCsgXp3e8X0PtK57QKnPtmktSX9
27lr7/BJ1/8k1u4rMYqsm+3J/NyZkrEB3ehEuQokuX9COjxs1v3FMXVc2YNIHmvm
6vMtt0MNmaqN5Gk1A7dCJkgLMdLt9WqNShzqk82RtLQzOuoNVr7jpnsKntYolXij
xvXm4DtEgFZ1D+FnKpIok3GYsJOxbcy0wbHsDwj6tV8c0XBpfBddAHGruB3QykRH
xuuulKx6rhSfIudrvQlQqCsOus1VkQD3jXHNfmnE8bIbiYIcBd40NW7qeLpni7wG
908zk21gaq7+v9ziJd5iLwBOGbvY82azuU1FSflx955xM9D+2tH2I9rRDhU9M+Aq
SKPmejPCFnVKWgF/eoFYpwRRBTzzzL7XP2Vl3t8Shn2rvDp/xCD73RPbAgMBAAGj
UzBRMB0GA1UdDgQWBBS5tpwImIAUWd0+ZLAp7Q3qyDQWfzAfBgNVHSMEGDAWgBS5
tpwImIAUWd0+ZLAp7Q3qyDQWfzAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEB
CwUAA4ICAQB4+S5y+npFdqr4DMyivTXtfzlRkZAMjy3876nwDzIoymzOyO1gGP0l
T66jDGfxTBaX3Zp7G8e9FP6bI25kCeoynUPmKjD2YPY3H+q7nb6q2HpgoAzTwtPg
Af7epg9rHASI/ghoKtoUjhrlzfbKmtAsbfIE/KrnTWoPmE2Qv8xqSyydNrDrBufd
jtBur2Z8Iq1uUcOtDXTMfBAUm+wXFrTmfr/+bH0EanK7hV9p9xqR4HXUjjC4+2zm
2Ghu0SlqrhfmoN63zXflPpgYIPQStPrOj0kWAMdVhya/QTM5dzIhvIKoSMmcvjEI
ca519fhs4TtyDYtNPd0SvGDkArDTwRthEYY1Ah7xCTsAdqHdQMdasecCiskiI+MV
7jHIT8goEdu6mE33rV9p3NieDwDnACx8NhAsJEj4V8OhjYdRUdx+XK9NtFa4kbMh
MEHy8BYnGS1H8qc5k2q4CPbB2Wuyfw5DO5gwkXrBptASZE8AnePBcqH091kiFSZd
1nJPQNTrAlvJ1w6IWKDVezd0bi4OHldYyrhVYGLDl9GXtbgQXJ24pzcGoHXFxDyi
1R61GusVOUaFSJKWhV3/uNNbiOBpqjXWGK10LdW5iU2UUipqHKLYXxUXRAuZOqxf
ZzzhvhxQBSSQBfIGNH166CsKcA8toxh+uwXRonzWfoktZU6PYYk1JA==
-----END CERTIFICATE-----`)

var CA_KEY = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

var GoproxyCa, goproxyCaErr = tls.X509KeyPair(CA_CERT, CA_KEY)
