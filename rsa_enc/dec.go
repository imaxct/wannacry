package rsa_enc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func Decrypt(data []byte) ([]byte, error) {
	privateKey := []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCJWoxJQn1JmwhGMsSGFFPjFtwe6ZiDB8TZOQnJ71/Ch3tISbCB
CpyLUgj72XPPTC6DEQzWdNrAH0cLFD8W/5Irt3ACaMNC8I9SxNqqYEDEmxqY/aHq
mBCgUmstolD90+Tyn3ySw7459D1zNTo1MoJrfu5bexINJ4eUICEHw8YUWwIDAQAB
AoGALJr2QfhQva2WM03bWnuRfob29yb/O1YzjCOk1Spim8bpt2EO6+kpS2lZt4g3
vtNLq47G74JFY+0EYkmx72MR8Zp1G6bxD/S5D5RxtsvODN8DBvYdzOYovLt71Zgf
FBwj5hB2IumdWVhgBAGJOz0aAlQ/IKU3aIKxQ8ffCguO98ECQQDjQJsvHngQdlHG
yDqBQY7GhFHBfX4cdnuegpef9ZYij2hd9PzXrbcC8kHbjAUVP2CWQeYtf6S9dzeT
dCeltz8hAkEAmrqkdgeVMf+rILjlFNZ5W2PqbytmaLKolEqIOafLZLBvogHA3c7G
4dU50t0Nn/u7suzlQlAM8wWQ/do8SKBP+wJBAJQ4Wj8kV1kdYv8NL6OIl9ABE7Xo
3O1BliVvted97buC36aQmK0vv3MrgSrqK5KNMLkwKCo6278718LT+twKJiECQDkd
rJflNK3AMuthVS7b/PA/ccqXurXrPU+AM8kUp4ADoTGsdAVszv2OOEoeT2k/A4qI
6BgSdLVA2MTcSmSEyQECQQDJT2BIxkVbMdibKks25naTYWknNWG1UoPAXpu/eA0O
izZe46FTnahzkfe7/SxzDkgonk6SuPiVKNJb09E27lul
-----END RSA PRIVATE KEY-----`)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}
