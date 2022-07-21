package utils

import (
	"errors"
	"github.com/hail-pas/GinStartKit/global/constant"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	if password == "" {
		return "", errors.New(constant.MessagePasswordNull)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

func VerifyHashAndPassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err == nil {
		return true
	}
	return false
}
