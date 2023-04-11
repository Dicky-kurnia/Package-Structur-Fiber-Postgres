package test

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
	"go-fiber-postgres/config"
	"go-fiber-postgres/helper"
	"go-fiber-postgres/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"testing"
)

var DB *gorm.DB
var defaultPasssword = "testing123"

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := godotenv.Load("../.env")
	if err != nil {
		helper.IsShouldPanic(err)
	}
	DB = config.NewFMCGPostgresDB()
}

func SelectRandomRows[T any](tableName string) (result T) {
	err := DB.Raw("SELECT * FROM " + tableName + " ORDER BY random() LIMIT 1").Scan(&result).Error
	if err != nil {
		panic(err)
	}
	return result
}

func TestCreateDummySales(t *testing.T) {
	pass, err := bcrypt.GenerateFromPassword([]byte(defaultPasssword), 10)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pass))

	sales := model.Sales{}
	err = faker.FakeData(&sales)
	if err != nil {
		panic(err)
	}
	sales.Username = faker.Username()
	sales.Name = faker.Name()
	sales.Password = string(pass)

	err = DB.Create(&sales).Error
	if err != nil {
		panic(err)
	}
}
