package pem

import (
	"bytes"
	"encoding/pem"
)

type Pem struct {
}

// Encode pem编码
func (p *Pem) Encode(Type string, Bytes []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := pem.Encode(&buf, &pem.Block{Type: Type, Bytes: Bytes}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode pem解码
func (p *Pem) Decode(data []byte) []byte {
	block, _ := pem.Decode(data)
	return block.Bytes
}
