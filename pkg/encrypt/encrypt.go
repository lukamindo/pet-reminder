package encrypt

import (
	"crypto/md5"

	"golang.org/x/crypto/bcrypt"
)

// Password encrypts password
func Password(pwd string) (string, error) {
	salt := []byte("ckUATqoW03mJvaRwNfbT3fu5I8mylUrK")
	h := md5.New()
	h.Write([]byte(pwd))
	md5Pass := h.Sum(salt)
	hash, err := bcrypt.GenerateFromPassword(md5Pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword returns boolean if pwd and hashedpwd are same
func CheckPassword(hashedPwd string, plainPwd string) bool {
	salt := []byte("ckUATqoW03mJvaRwNfbT3fu5I8mylUrK")
	h := md5.New()
	h.Write([]byte(plainPwd))
	md5Pass := h.Sum(salt)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, md5Pass)
	if err != nil {
		return false
	}
	return true
}
