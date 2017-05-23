package rsa_enc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func Encrypt(data []byte) ([]byte, error) {
	publicKey := []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCJWoxJQn1JmwhGMsSGFFPjFtwe
6ZiDB8TZOQnJ71/Ch3tISbCBCpyLUgj72XPPTC6DEQzWdNrAH0cLFD8W/5Irt3AC
aMNC8I9SxNqqYEDEmxqY/aHqmBCgUmstolD90+Tyn3ySw7459D1zNTo1MoJrfu5b
exINJ4eUICEHw8YUWwIDAQAB
-----END PUBLIC KEY-----`)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pubInterface.(*rsa.PublicKey), data)
}
