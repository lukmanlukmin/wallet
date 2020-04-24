package helper

import (
	"fmt"
	"log"

	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashString(payload string) string {
	// bytes, err := bcrypt.GenerateFromPassword([]byte(payload), 14)
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)
}

func CompareHash(plainText string, chipperText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(plainText), []byte(chipperText))
	return err == nil
}

var jwtKey = []byte("privytestsecret")

type Claims struct {
	UserID   int    `json:"userId"`
	UserType string `json:"usertype"`
	jwt.StandardClaims
}

func CreateToken(userID int, userType string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (Claims, int, error) {
	claims := &Claims{}
	var status = 0
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// unauthorized
			fmt.Println("signature invalid")
			status = 1
		}
		// bad request
		status = 2
	}
	if !tkn.Valid {
		// unauthorized
		fmt.Println("token invalid")
		status = 1
	}
	return *claims, status, err
}
