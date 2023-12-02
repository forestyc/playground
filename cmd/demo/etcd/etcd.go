package main

import (
	"fmt"
	"github.com/forestyc/playground/cmd/demo/pki/ca"
)

var storage *ca.Storage
var err error

func main() {
	storage, err = ca.NewStorage([]string{
		"81.70.188.168:12379",
		"140.143.163.171:12379",
		"101.42.23.168:12379",
	})
	if err != nil {
		panic(err)
	}
	defer storage.Close()
	//var rootCsr = x509.Certificate{
	//	Version:      3,
	//	SerialNumber: big.NewInt(time.Now().Unix()),
	//	Subject: pkix.Name{
	//		Country:            []string{"CN"},
	//		Province:           []string{"Shanghai"},
	//		Locality:           []string{"Shanghai"},
	//		Organization:       []string{"JediLtd"},
	//		OrganizationalUnit: []string{"JediProxy"},
	//		CommonName:         "Jedi Root CA",
	//	},
	//	NotBefore:             time.Now(),
	//	NotAfter:              time.Now().AddDate(10, 0, 0),
	//	BasicConstraintsValid: true,
	//	IsCA:                  true,
	//	MaxPathLen:            1,
	//	MaxPathLenZero:        false,
	//	KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	//}

	//if err := storage.Put("root-test", rootCsr); err != nil {
	//	panic(err)
	//}
	cert, err := storage.Get("root-test")
	if err != nil {
		panic(err)
	}
	fmt.Println(cert)
}
