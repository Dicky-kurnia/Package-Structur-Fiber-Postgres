package model

type Sales struct {
	Id       uint   `gorm:"primary_key;autoIncrement" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
