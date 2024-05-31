package ca

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"github.com/forestyc/playground/pkg/encoding/pem"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	randM "math/rand"
	"time"
)

type CA struct {
	storage *Storage
	pem     pem.Pem
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

// CreateCaCertificate create middle certificate and private key.
// Returns certPem, privatePem and error.
func (ca *CA) CreateCaCertificate(bit int, cn string, root bool, option ...Option) ([]byte, []byte, error) {
	var certRoot *x509.Certificate
	var privRoot *rsa.PrivateKey
	var err error
	if !root {
		// get root certificate
		certRoot, privRoot, err = ca.getCertAndKey(PrefixRoot)
		if err != nil {
			return nil, nil, err
		}
	}

	// create certificate
	option = append(option, WithIsCa(true))
	cert, priv, err := ca.createCertificate(certRoot, privRoot, bit, cn, option...)
	if err != nil {
		return nil, nil, err
	}

	// store to etcd
	if err = ca.store(PrefixMiddle+cn, cert, priv); err != nil {
		return nil, nil, err
	}
	return cert, priv, err
}

// CreateServerCertificate create terminal certificate and private key.
// Returns certPem, privatePem and error.
func (ca *CA) CreateServerCertificate(bit int, cn string, option ...Option) ([]byte, []byte, error) {
	// get parent certificate
	certParent, privParent, err := ca.getCertAndKey(PrefixMiddle)
	if err != nil {
		return nil, nil, err
	}

	// create certificate
	option = append(option, WithIsCa(false), WithMaxPathLen(0))
	cert, priv, err := ca.createCertificate(certParent, privParent, bit, cn, option...)
	if err != nil {
		return nil, nil, err
	}

	// store to etcd
	if err = ca.store(PrefixTerminal+cn, cert, priv); err != nil {
		return nil, nil, err
	}
	return cert, priv, err
}

// createCertificate create certificate and private key.
func (ca *CA) createCertificate(parentCert *x509.Certificate, parentKey *rsa.PrivateKey, bit int, cn string, option ...Option) ([]byte, []byte, error) {
	// generate rsa key
	priv, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return nil, nil, err
	}

	// create x509.Certificate
	tpl := NewCertificate(cn, option...)
	if parentCert == nil {
		parentCert = &tpl.Certificate
	}
	if parentKey == nil {
		parentKey = priv
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tpl.Certificate, parentCert, &priv.PublicKey, parentKey)
	if err != nil {
		return nil, nil, err
	}

	// pem encode
	privPem, err := ca.pem.Encode("RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(priv))
	if err != nil {
		return nil, nil, err
	}
	certPem, err := ca.pem.Encode("CERTIFICATE", cert)
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

// getMiddleCertAndKey get and parse certificate and private key of intermediate
func (ca *CA) getCertAndKey(prefix string) (*x509.Certificate, *rsa.PrivateKey, error) {
	var err error
	// get certificate
	kvs, err := ca.storage.Get(prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, nil, err
	} else if len(kvs) == 0 {
		return nil, nil, errors.New("no root certificate!")
	}

	// pick one
	randM.New(randM.NewSource(time.Now().UnixNano()))
	idx := randM.Intn(len(kvs))

	// json unmarshal
	kv := KeyPair{}
	if err = json.Unmarshal(kvs[idx].Value, &kv); err != nil {
		return nil, nil, err
	}

	return ca.parseCertAndKey(kv.Cert, kv.Key)
}

// parseCertAndKey parse certificate and private key formatted by PEM
func (ca *CA) parseCertAndKey(certPem, keyPem []byte) (*x509.Certificate, *rsa.PrivateKey, error) {
	cert, err := x509.ParseCertificate(ca.pem.Decode(certPem))
	if err != nil {
		return nil, nil, err
	}
	priv, err := x509.ParsePKCS1PrivateKey(ca.pem.Decode(keyPem))
	if err != nil {
		return nil, nil, err
	}
	return cert, priv, nil
}
