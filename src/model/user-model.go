package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type UserModelInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetJSONValue() (string, error)
	SetId(string)
	EncryptPassword()
}

type userModel struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

func NewUserModel(email, password, name string, age int8) *userModel {
	return &userModel{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}
func (userModel *userModel) SetId(id string) {
	userModel.Id = id
}

func (userModel *userModel) GetJSONValue() (string, error) {
	json, err := json.Marshal(userModel)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (userModel *userModel) GetEmail() string {
	return userModel.Email
}
func (userModel *userModel) GetPassword() string {
	return userModel.Password
}
func (userModel *userModel) GetAge() int8 {
	return userModel.Age
}
func (userModel *userModel) GetName() string {
	return userModel.Name
}

func (userModel *userModel) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userModel.Password))
	userModel.Password = hex.EncodeToString(hash.Sum(nil))
}
