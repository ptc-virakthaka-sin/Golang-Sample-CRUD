package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
	return string(bytes), err
}

func Verify(text, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}
