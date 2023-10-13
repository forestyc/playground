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
	Cert []byte
	Key  []byte
}

// NewCA new ca
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
func (ca *CA) CreateRootCertificate(bit int, cn string, option ...Option) ([]byte, []byte, error) {
	// create certificate
	option = append(option, WithIsCa(true))
	cert, priv, err := ca.createCertificate(nil, bit, cn, option...)
	if err != nil {
		return nil, nil, err
	}

	// store to etcd
	if err = ca.store(PrefixRoot+cn, cert, priv); err != nil {
		return nil, nil, err
	}
	return cert, priv, err
}

// CreateMiddleCertificate create middle certificate and private key.
// Returns certPem, privatePem and error.
func (ca *CA) CreateMiddleCertificate(parent *x509.Certificate, bit int, cn string, option ...Option) ([]byte, []byte, error) {
	// create certificate
	option = append(option, WithIsCa(true))
	cert, priv, err := ca.createCertificate(parent, bit, cn, option...)
	if err != nil {
		return nil, nil, err
	}

	// store to etcd
	if err = ca.store(PrefixMiddle+cn, cert, priv); err != nil {
		return nil, nil, err
	}
	return cert, priv, err
}

// CreateTerminalCertificate create terminal certificate and private key.
// Returns certPem, privatePem and error.
func (ca *CA) CreateTerminalCertificate(parent *x509.Certificate, bit int, cn string, option ...Option) ([]byte, []byte, error) {
	// create certificate
	option = append(option, WithIsCa(false), WithMaxPathLen(0))
	cert, priv, err := ca.createCertificate(parent, bit, cn, option...)
	if err != nil {
		return nil, nil, err
	}

	// store to etcd
	if err = ca.store(PrefixMiddle+cn, cert, priv); err != nil {
		return nil, nil, err
	}
	return cert, priv, err
}

// createCertificate create certificate and private key.
func (ca *CA) createCertificate(parent *x509.Certificate, bit int, cn string, option ...Option) ([]byte, []byte, error) {
	// generate rsa key
	priv, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return nil, nil, err
	}

	// create x509.Certificate
	tpl := NewCertificate(cn, option...)
	if parent == nil {
		parent = &tpl.Certificate
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tpl.Certificate, parent, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	// pem encode
	privPem, err := encoding.PemEncode("RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(priv))
	if err != nil {
		return nil, nil, err
	}
	certPem, err := encoding.PemEncode("CERTIFICATE", cert)
	if err != nil {
		return nil, nil, err
	}
	return certPem, privPem, nil
}

// store to etcd
func (ca *CA) store(name string, certPem, privPem []byte) error {
	pair := KeyPair{
		Cert: certPem,
		Key:  privPem,
	}
	jsonByte, err := json.Marshal(&pair)
	if err != nil {
		return err
	}
	if err = ca.storage.Put(name, string(jsonByte)); err != nil {
		return err
	}
	return nil
}
