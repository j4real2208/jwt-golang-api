package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey =[]byte("mysupersecretkey")


func GenrateJWT() (string,error)  {
	token := jwt.New(jwt.SigningMethodHS256)

	claims :=  token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Jojo Johnson"
	claims["times"] = time.Now().Add(time.Hour).Unix()

	tokenString , err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString , nil
}


func main() {
	fmt.Println("Hello World ")

	tokenString , err := GenrateJWT()
	if err != nil {
		log.Println(err)
	}
	println(tokenString)
}