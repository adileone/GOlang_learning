package models

type User struct {
	Id       int
	Username string `orm:"size(100);unique"`
	Password string `orm:"size(100);unique"`
	Email    string `orm:"size(64);unique"`
	Balance  float32
}

func (u *User) TableName() string {
	return "users"
}
