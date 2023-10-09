package crypto

import "encoding/base64"

// Base64Encode base64编码
func Base64Encode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64Decode base64编码
func Base64Decode(data string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(data)
}
