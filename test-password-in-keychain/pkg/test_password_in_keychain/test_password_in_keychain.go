package test_password_in_keychain

import (
	"fmt"
	"log"

	"github.com/zalando/go-keyring"
)

const (
	Service = "test-password-in-keychain"
	User    = "example"
)

func SetPassword(password string) {
	err := keyring.Set(Service, User, password)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPassword() (string, error) {
	return keyring.Get(Service, User)
}

func PrintPassword() {
	password, err := GetPassword()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(password)
}
