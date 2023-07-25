package model

type userModel struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (userModel *userModel) SetId(id string) {
	userModel.id = id
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

func (userModel *userModel) GetId() string {
	return userModel.id
}
