package encryption

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)


func Hashing(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", err
	}    // GenerateFromPassword returns a byte slice so we need to
	return string(hash), nil
}

func CompareHash(hashedStr string, str string) (error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(str)); err != nil {
		return err
	}
	return nil
}
