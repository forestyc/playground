package utils

import (
	"bytes"
	"encoding/binary"
)

func ToBytes(obj any) ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func FromBytes(data []byte, obj any) error {
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, obj)
	if err != nil {
		return err
	}
	return err
}
