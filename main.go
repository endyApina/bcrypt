package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	defaultString := "apinaendy@gmail.com"

	hashedString, err := encryptString(defaultString)
	checkErr(err)

	_, err = decryptStrings(hashedString, "apinaendy@gmail.om")

	log.Println(hashedString)
}

//encryptString encrypts the string
func encryptString(text string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(text), 14)
	checkErr(err)
	hashString := string(byte)
	return hashString, nil
}

//decryptStrings decrypts a string envrypted by bcrypt
func decryptStrings(encryptedString string, initialString string) (bool, error) {
	status := compareStrings(encryptedString, initialString)
	log.Println(status)
	return status, nil
}

//compareStrings compares 2 string and returns a bool
func compareStrings(encryptedString string, initialString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedString), []byte(initialString))
	checkErr(err)
	return err == nil
}

//checkErr checks and panic when there is an error.
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
