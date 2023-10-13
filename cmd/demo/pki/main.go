package main

import (
	"os"

	"github.com/forestyc/playground/pkg/security/pki"
)

func main() {
	ca, err := pki.NewCA([]string{
		"81.70.188.168:12379",
		"140.143.163.171:12379",
		"101.42.23.168:12379",
	})
	if err != nil {
		panic(err)
	}
	rootCert, rootPrive, err := ca.CreateRootCertificate(2048, "forestyc CA")
	if err != nil {
		panic(err)
	}
	f1, err := os.Create("root.pem")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f1.Write([]byte(rootCert))
	f2, err := os.Create("root.key")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	f2.Write([]byte(rootPrive))
}
