package pki

import (
	"crypto/x509/pkix"
	"math/big"
	"time"
)

type Certificate struct {
	Version               int       `json:"Version,omitempty"`
	SerialNumber          *big.Int  `json:"SerialNumber,omitempty"`
	Issuer                pkix.Name `json:"Issuer,omitempty"`
	Subject               pkix.Name `json:"Subject,omitempty"`
	NotBefore             time.Time `json:"NotBefore,omitempty"`
	NotAfter              time.Time `json:"NotAfter,omitempty"`
	KeyUsage              int       `json:"KeyUsage,omitempty"`
	BasicConstraintsValid bool      `json:"BasicConstraintsValid,omitempty"`
	IsCA                  bool      `json:"IsCA,omitempty"`
	MaxPathLen            int       `json:"MaxPathLen,omitempty"`
	MaxPathLenZero        bool      `json:"MaxPathLenZero,omitempty"`
}
