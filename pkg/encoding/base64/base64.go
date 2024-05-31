package base64

import (
	"encoding/base64"
	"strings"
)

const (
	plus  = "+"
	slash = "/"
)

type Encoding int

const (
	UrlEncoding Encoding = iota + 1
	StdEncoding
)

type Base64 struct {
	encoding Encoding
}

// Encode returns the URL-safe base64 encoding of src.
func (b *Base64) Encode(data []byte) string {
	switch b.encoding {
	case UrlEncoding:
		return base64.URLEncoding.EncodeToString(data)
	case StdEncoding:
		return base64.StdEncoding.EncodeToString(data)
	default:
		return base64.URLEncoding.EncodeToString(data)
	}
}

// Decode returns the bytes represented by the base64 string data
func (b *Base64) Decode(data string) ([]byte, error) {
	if strings.Contains(data, plus) || strings.Contains(data, slash) {
		if strings.ContainsRune(data, base64.StdPadding) {
			return base64.StdEncoding.DecodeString(data)
		} else {
			return base64.RawStdEncoding.DecodeString(data)
		}
	} else {
		if strings.ContainsRune(data, base64.StdPadding) {
			return base64.URLEncoding.DecodeString(data)
		} else {
			return base64.RawURLEncoding.DecodeString(data)
		}
	}
}

// WithEncoding uses encoding
func (b *Base64) WithEncoding(encoding Encoding) *Base64 {
	b.encoding = encoding
	return b
}
