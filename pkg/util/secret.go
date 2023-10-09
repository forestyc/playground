package util

import (
	"strconv"
	"strings"
	"time"

	"github.com/forestyc/playground/pkg/crypto"
	uuid "github.com/satori/go.uuid"
)

// GenAppKey 生成appKey
func GenAppKey() string {
	return genUuid()
}

// GenAppSecret 生成appSecret
// sm3(appKey + name + timestamp + random)
func GenAppSecret(appKey, name string) (string, error) {
	nano := time.Now().UnixNano()
	data := appKey + name + strconv.FormatInt(nano, 10) + genUuid()
	sm3 := crypto.SM3{}
	appSecret, err := sm3.Sum([]byte(data))
	if err != nil {
		return "", err
	}
	return appSecret, err
}

func genUuid() string {
	appKey := uuid.NewV4().String()
	return strings.Replace(appKey, "-", "", -1)
}
