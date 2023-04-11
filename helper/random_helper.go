package helper

import (
	"github.com/google/uuid"
	"regexp"
)

func GenRandomInt() string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		panic(err)
	}

	str := reg.ReplaceAllString(uuid.New().String(), "")
	return str[:len(str)/2]
}
