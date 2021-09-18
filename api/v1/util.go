package v1

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
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

func randomCode() string {
	c := 3
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	var builder strings.Builder
	for i := 0; i < len(b); i++ {
		builder.WriteString(fmt.Sprintf("%03d", b[i]))
	}
	return builder.String()
}
