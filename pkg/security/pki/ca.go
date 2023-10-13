package pki

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"github.com/forestyc/playground/pkg/encoding"
)

type CA struct {
	storage *Storage
}

type KeyPair struct {
	Cert string
	Key  string
}

// NewCA
func NewCA(endpoints []string) (*CA, error) {
	var err error
	pki := &CA{}
	if pki.storage, err = NewStorage(endpoints); err != nil {
		return nil, err
	}
	return pki, nil
}

// CreateRootCertificate create root certificate and private key.
// Returns certPem, privatePem and error.
func (ca CA) CreateRootCertificate(bit int, cn string, option ...Option) (string, string, error) {
	tpl := NewCertificate(cn, option...)
	priv, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return "", "", err
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tpl.Certificate, &tpl.Certificate, &priv.PublicKey, priv)
	if err != nil {
		return "", "", err
	}
	// private key pem
	privPem, err := encoding.PemEncode("RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(priv))
	if err != nil {
		return "", "", err
	}
	certPem, err := encoding.PemEncode("CERTIFICATE", cert)
	if err != nil {
		return "", "", err
	}
	// storage
	pair := KeyPair{
		Cert: certPem,
		Key:  privPem,
	}
	jsonByte, err := json.Marshal(&pair)
	if err != nil {
		return "", "", err
	}
	if err = ca.storage.Put(PrefixRoot+cn, string(jsonByte)); err != nil {
		return "", "", err
	}
	return certPem, privPem, nil
}
