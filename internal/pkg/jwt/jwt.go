package jwt

import (
	"errors"
	"time"

	"github.com/vanbien2402/first-web-demo/internal/pkg/constant"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

//Claim struct to generate jwt
type Claim struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

//GenerateJWT generate jwt by userName and email
func GenerateJWT(userName, email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(constant.JWTExpiration * time.Hour)
	claims := Claim{
		UserName: userName,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

//ValidateToken validate token
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("could not parse claim")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
