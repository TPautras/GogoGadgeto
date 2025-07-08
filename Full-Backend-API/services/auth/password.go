package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func ComparePasswords(hashed string, plain []byte) bool {
	fmt.Printf("Comparing hashed password: %s with plain password: %s\n", hashed, string(plain))
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)

	if err != nil {
		fmt.Printf("Password comparison failed: %v\n", err)
		return false
	}
	fmt.Printf("Password comparison success: %v\n", err)
	return true
}
