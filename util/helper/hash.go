package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashString(payload string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(password)
}

func CompareHash(origin string, hashed string) bool {
	hashedByte := []byte(hashed)
	originByte := []byte(origin)
	err := bcrypt.CompareHashAndPassword(hashedByte, originByte)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
