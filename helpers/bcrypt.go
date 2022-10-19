package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(str string) string {
	salt := 8
	byteToHash := []byte(str)
	hash, err := bcrypt.GenerateFromPassword(byteToHash, salt)
	if err != nil {
		panic("Failed to hash string")
	}

	return string(hash)
}

func ComparePass(hash, valueToCompare string) bool {
	hashByte, valueByte := []byte(hash), []byte(valueToCompare)

	err := bcrypt.CompareHashAndPassword(hashByte, valueByte)

	return err == nil
}
