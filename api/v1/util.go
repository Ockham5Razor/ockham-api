package v1

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func checkEncrypt(hashed string, toCheckRawString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(toCheckRawString))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func encrypt(rawString string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawString), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

