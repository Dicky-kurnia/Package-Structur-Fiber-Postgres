package model

type BranchOffice struct {
	Id   uint   `gorm:"primary_key;autoIncrement" json:"id"`
	Name string `json:"name"`
}
