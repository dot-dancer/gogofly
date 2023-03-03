package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var stSigningKey = []byte(viper.GetString("jwt.signingKey"))

type JwtCustClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id uint, name string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString(stSigningKey)
}

func ParseToken(tokenStr string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return stSigningKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("Invalid Token")
	}

	return iJwtCustClaims, err
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}

	return true
}
