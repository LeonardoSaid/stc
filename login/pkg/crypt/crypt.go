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

func CompareHash(hash string, secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
}
