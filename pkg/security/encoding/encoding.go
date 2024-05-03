package encoding

import (
	"bytes"
	"encoding/base64"
	"encoding/pem"
)

// Base64Encode base64编码
func Base64Encode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64Decode base64解码
func Base64Decode(data string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(data)
}

// PemEncode pem编码
func PemEncode(Type string, Bytes []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := pem.Encode(&buf, &pem.Block{Type: Type, Bytes: Bytes}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// PemDecode pem解码
func PemDecode(data []byte) []byte {
	block, _ := pem.Decode(data)
	return block.Bytes
}
