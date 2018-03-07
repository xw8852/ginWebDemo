package util

import "regexp"

func ValidatePhone(phone string) bool {
	m, e := regexp.MatchString("1[3|5|6|7|8|9][0-9]{9}", phone)
	if e != nil {
		return false
	}
	return m;
}
