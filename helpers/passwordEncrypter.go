package helpers

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EncryptPassword(password string) (string, error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", errors.New("error while generating bcrypt password")
	}

	return string(encryptedPass), nil
}

// IsPasswordCorrect decrypts & verifies if the password is correct.
func IsPasswordCorrect(hashedPassword, existingPlainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(existingPlainPassword))
	if err != nil {
		log.Println(err)
		return false, errors.New("wrong password")
	}

	return true, nil
}
