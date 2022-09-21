package jwt

import (
	"errors"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"

	"github.com/vanbien2402/first-web-demo/internal/pkg/constant"
)

var jwtKey = []byte("supersecretkey")

//Claim struct to generate jwt
type Claim struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	jwtGo.StandardClaims
}

//GenerateJWT generate jwt by userName and email
func GenerateJWT(userName, email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(constant.JWTExpiration * time.Hour)
	claims := Claim{
		UserName: userName,
		Email:    email,
		StandardClaims: jwtGo.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

//ValidateToken validate token
func ValidateToken(signedToken string) (err error) {
	token, err := jwtGo.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwtGo.Token) (interface{}, error) {
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
