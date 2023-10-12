package pki

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
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
func (ca CA) CreateRootCertificate(bit int, cn string, option ...Option) ([]byte, []byte, error) {
	tpl := NewCertificate(cn, option...)
	priv, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return nil, nil, err
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tpl.Certificate, &tpl.Certificate, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}
	// private key pem
	var bufPriv, bufCert bytes.Buffer
	if err := pem.Encode(&bufPriv, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}); err != nil {
		return nil, nil, err
	}
	if err := pem.Encode(&bufCert, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		return nil, nil, err
	}
	// storage
	pair := KeyPair{
		Cert: bufCert.String(),
		Key:  bufPriv.String(),
	}
	jsonByte, err := json.Marshal(&pair)
	if err != nil {
		return nil, nil, err
	}
	if err = ca.storage.Put(PrefixRoot+cn, string(jsonByte)); err != nil {
		return nil, nil, err
	}
	return bufCert.Bytes(), bufPriv.Bytes(), nil
}
