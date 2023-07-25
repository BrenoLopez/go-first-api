package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (userModel *userModel) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userModel.password))
	userModel.password = hex.EncodeToString(hash.Sum(nil))
}
