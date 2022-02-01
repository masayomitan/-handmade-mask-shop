package component

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"math/rand"
)


func HashPassword(pass string) (string) {
	storePass := pass
	hash, err := bcrypt.GenerateFromPassword([]byte(storePass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
  }
	return string(hash);
}


func CheckPassword(hashedPassword string, password string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//numに指定文字数を入れる
func RandString(num int) (string) {
	const rs1Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, num)
	for i := range b {
			b[i] =  rs1Letters[int(rand.Int63()%int64(len(rs1Letters)))]
	}
	return string(b)
}