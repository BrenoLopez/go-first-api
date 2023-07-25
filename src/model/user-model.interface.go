package model

type UserModelInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	SetId(string)
	EncryptPassword()
}

func NewUserModel(email, password, name string, age int8) *userModel {
	return &userModel{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}
