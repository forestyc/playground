package ca

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

type Option func(*Certificate)

type Certificate struct {
	x509.Certificate
}

// NewCertificate new certificate with CommonName.
// To Add other information, use WithXXX().
func NewCertificate(cn string, option ...Option) Certificate {
	c := Certificate{
		Certificate: x509.Certificate{
			SerialNumber: big.NewInt(time.Now().UnixNano()),
			Subject: pkix.Name{
				CommonName: cn,
			},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().AddDate(1, 0, 0),
			BasicConstraintsValid: true,
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
			IssuingCertificateURL: nil,
		},
	}
	for _, o := range option {
		o(&c)
	}
	return c
}

// WithSubjectCountry add country for Subject
func WithSubjectCountry(country []string) Option {
	return func(c *Certificate) {
		c.Subject.Country = country
	}
}

// WithSubjectProvince add province for Subject
func WithSubjectProvince(province []string) Option {
	return func(c *Certificate) {
		c.Subject.Province = province
	}
}

// WithSubjectLocality add locality for Subject
func WithSubjectLocality(locality []string) Option {
	return func(c *Certificate) {
		c.Subject.Locality = locality
	}
}

// WithSubjectOrganization add organization for Subject
func WithSubjectOrganization(organization []string) Option {
	return func(c *Certificate) {
		c.Subject.Organization = organization
	}
}

// WithSubjectOrganizationalUnit add organizationalUnit for Subject
func WithSubjectOrganizationalUnit(organizationalUnit []string) Option {
	return func(c *Certificate) {
		c.Subject.OrganizationalUnit = organizationalUnit
	}
}

// WithMaxPathLen add max length for certificate link
func WithMaxPathLen(len int) Option {
	return func(c *Certificate) {
		c.MaxPathLen = len
	}
}

// WithIsCa set IsCa
func WithIsCa(isCa bool) Option {
	return func(c *Certificate) {
		c.IsCA = isCa
	}
}

// WithNotAfter set NotAfter
func WithNotAfter(t time.Time) Option {
	return func(c *Certificate) {
		c.NotAfter = t
	}
}
