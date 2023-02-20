package model

import (
	"github.com/Baal19905/playground/go-zero/epidemic/pkg/crypto"
	"gorm.io/gorm"
)

func (user *EpidemicUser) AfterFind(tx *gorm.DB) (err error) {
	cryptoInstance := crypto.AES{}
	if user.Name = cryptoInstance.DecryptWithBase64(string(user.Name), config.SymmetricKey); err != nil {
		return
	}
	if user.ResidentialAddressCode = cryptoInstance.DecryptWithBase64(string(user.ResidentialAddressCode), config.SymmetricKey); err != nil {
		return
	}
	if user.ResidentialAddressDetail = cryptoInstance.DecryptWithBase64(string(user.ResidentialAddressDetail), config.SymmetricKey); err != nil {
		return
	}
	if user.ResidentialAddressName = cryptoInstance.DecryptWithBase64(string(user.ResidentialAddressName), config.SymmetricKey); err != nil {
		return
	}
	return
}
