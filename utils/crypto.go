package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(stText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(stText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func CompareHashAndPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
