package crypt

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"wood/pkg/constant"
)

func GeneratePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, constant.PasswordHashBytes)
	if err != nil {
		return "", errors.New("error: failed to generate password hash")
	}

	return fmt.Sprintf("%x", h), nil
}

func PasswordCrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyCryptPwd(password string, loginPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginPwd)) //验证（对比）
	if err != nil {
		return false
	} else {
		return true
	}
}
