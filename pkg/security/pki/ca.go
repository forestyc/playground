package pki

import (
	"crypto/rand"
	"crypto/x509"
)

type Pki struct {
}

func NewPki() Pki {
	return Pki{}
}

func (pki Pki) CreateCertificate() {
	x509.CreateCertificate(rand.Reader)
}
