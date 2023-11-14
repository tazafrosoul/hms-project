package utility

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func Verify(password, hashed_pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_pw), []byte(password))
	return err == nil
}
