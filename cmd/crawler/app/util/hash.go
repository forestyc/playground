package util

import "hash/crc32"

// Sum 取字符串哈希值
func Sum(data string) int {
	v := int(crc32.ChecksumIEEE([]byte(data)))
	if v >= 0 {
		return v
	} else if v <= 0 {
		return -v
	}
	return 0
}
