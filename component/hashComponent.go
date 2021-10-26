package component

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)


func HashPassword (pass string) string {
	storePass := "pass"
	hash, err := bcrypt.GenerateFromPassword([]byte(storePass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
}
	return string(hash);
}