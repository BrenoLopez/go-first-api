package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserModelInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	EncryptPassword()
}

type userModel struct {
	email    string
	password string
	name     string
	age      int8
}

func NewUserModel(email, password, name string, age int8) *userModel {
	return &userModel{
		email,
		password,
		name,
		age,
	}
}

func (userModel *userModel) GetEmail() string {
	return userModel.email
}
func (userModel *userModel) GetPassword() string {
	return userModel.password
}
func (userModel *userModel) GetAge() int8 {
	return userModel.age
}
func (userModel *userModel) GetName() string {
	return userModel.name
}

func (userModel *userModel) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userModel.password))
	userModel.password = hex.EncodeToString(hash.Sum(nil))
}
