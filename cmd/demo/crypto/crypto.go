package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/forestyc/playground/pkg/security/crypto"
)

func main() {
	rsa := crypto.RSA{}

	priv, pub, err := rsa.GenerateKey()
	if err != nil {
		panic(err)
	}
	plaintext := "daihouda"
	sum := md5.Sum([]byte(plaintext))
	fmt.Println(hex.EncodeToString(sum[0:]))
	sign, err := rsa.SignWithBase64(sum[0:], priv)
	if err != nil {
		panic(err)
	}
	fmt.Println("sign:", string(sign))
	success, err := rsa.VerifyWithBase64(sum[0:], pub, sign)
	if err != nil {
		panic(err)
	}
	if success == false {
		fmt.Println("failed")
	} else {
		fmt.Println("success")
	}
}
