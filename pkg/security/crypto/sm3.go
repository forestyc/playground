package crypto

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm3"
)

type SM3 struct {
}

// Sum 摘要
func (s SM3) Sum(data []byte) (string, error) {
	if len(data) == 0 {
		return "", errors.New(InvalidParameters)
	}
	h := sm3.New()
	h.Write([]byte(data))
	sumByte := h.Sum(nil)
	return hex.EncodeToString(sumByte), nil
}
