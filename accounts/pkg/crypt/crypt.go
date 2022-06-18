package crypt

import "golang.org/x/crypto/bcrypt"

func HashSecret(secret string) (string, error) {
	pass := []byte(secret)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
