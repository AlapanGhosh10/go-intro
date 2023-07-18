package helper

import "strings"

var roll = 0;

func IsValid(fname string, lname string, email string) (bool, bool) {
	isValidName := len(fname) > 2 && len(lname) > 2
	isValidEmail := strings.Contains(email, "@")
	return isValidName, isValidEmail
}

func GenerateRoll() int {
	roll++
	return roll
}