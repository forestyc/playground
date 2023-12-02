package utils

import "hash/crc32"

// HashString get a unique value from a string
func HashString(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}
