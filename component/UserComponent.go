package component

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword (pass string) string {

	storePass := pass
	hash, err := bcrypt.GenerateFromPassword([]byte(storePass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
  }
	return string(hash);
}

func CheckPassword (hashedPassword string, password string) (error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}