package main

import (
	"github.com/forestyc/playground/pkg/security/pki/ca"
	"os"
	"time"
)

func main() {
	CA, err := ca.NewCA([]string{
		"81.70.188.168:12379",
		"140.143.163.171:12379",
		"101.42.23.168:12379",
	})
	if err != nil {
		panic(err)
	}
	Root(CA)
	Middle(CA)
	Server(CA)
}

func Root(CA *ca.CA) []byte {
	cert, priv, err := CA.CreateCaCertificate(
		2048,
		"forestyc root",
		true,
		ca.WithNotAfter(time.Now().AddDate(10, 0, 0)),
	)
	if err != nil {
		panic(err)
	}
	f1, err := os.Create("root.pem")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f1.Write([]byte(cert))
	f2, err := os.Create("root.key")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	f2.Write([]byte(priv))
	return cert
}

func Middle(CA *ca.CA) []byte {
	cert, priv, err := CA.CreateCaCertificate(
		2048,
		"forestyc intermediate",
		false,
		ca.WithNotAfter(time.Now().AddDate(5, 0, 0)),
	)
	if err != nil {
		panic(err)
	}
	f1, err := os.Create("middle.pem")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f1.Write([]byte(cert))
	f2, err := os.Create("middle.key")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	f2.Write([]byte(priv))
	return cert
}

func Server(CA *ca.CA) []byte {
	cert, priv, err := CA.CreateServerCertificate(
		2048,
		"forestyc server",
	)
	if err != nil {
		panic(err)
	}
	f1, err := os.Create("terminal.pem")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f1.Write([]byte(cert))
	f2, err := os.Create("terminal.key")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	f2.Write([]byte(priv))
	return cert
}
