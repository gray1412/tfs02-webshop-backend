package validator

import (
	"regexp"
)

func CheckLength(data string, maxSize int) bool {
	return len(data) < maxSize
}
func CheckPhone(data string) bool {
	regex := "(84|0[3|5|7|8|9])+([0-9]{8})\\b"
	match, _ := regexp.MatchString(regex, data)
	return match
}
func CheckMail(data string) bool {
	if !CheckLength(data, 255) {
		return false
	}
	regex := "^[a-zA-Z][a-zA-Z0-9_\\.]{4,32}@[a-zA-Z0-9]{2,}(\\.[a-zA-Z0-9]{2,4}){1,2}$"
	match, _ := regexp.MatchString(regex, data)
	return match
}
func CheckName(data string) bool {
	regex := "^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$"
	match, _ := regexp.MatchString(regex, data)
	return match
}
