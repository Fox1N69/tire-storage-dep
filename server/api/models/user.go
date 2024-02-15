package models

type User struct {
	Id       string
	Name     string
	Email    string `gorm: "unique"`
	Password string
}
