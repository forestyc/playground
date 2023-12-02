package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/forestyc/playground/pkg/crypto"
	"github.com/google/uuid"
)

// GenAppKey 生成appKey
func GenAppKey() string {
	return GenUuid()
}

// GenAppSecret 生成appSecret
// sm3(appKey + name + timestamp + random)
func GenAppSecret(appKey, name string) (string, error) {
	nano := time.Now().UnixNano()
	data := appKey + name + strconv.FormatInt(nano, 10) + GenUuid()
	sm3 := crypto.SM3{}
	appSecret, err := sm3.Sum([]byte(data))
	if err != nil {
		return "", err
	}
	return appSecret, err
}

func GenUuid() string {
	appKey := uuid.New().String()
	return strings.Replace(appKey, "-", "", -1)
}
