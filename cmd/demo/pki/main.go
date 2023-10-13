package main

import (
	"crypto/x509"
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

	Root(ca)
	//middleCert := Middle(ca, rootCert)
	//Terminal(ca, middleCert)
}

func Root(ca *pki.CA) []byte {
	cert, priv, err := ca.CreateRootCertificate(
		2048,
		"forestyc CA",
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

func Middle(ca *pki.CA, parent *x509.Certificate) []byte {
	cert, priv, err := ca.CreateMiddleCertificate(
		parent,
		2048,
		"forestyc CA",
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

func Terminal(ca *pki.CA, parent *x509.Certificate) []byte {
	cert, priv, err := ca.CreateTerminalCertificate(
		parent,
		2048,
		"forestyc application",
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
