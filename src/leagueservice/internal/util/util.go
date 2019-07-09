package util

import "unicode"

func VerifyPassword(password string) bool {
	var sixToTwenty, number, upper, lower bool
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		}
	}

	sixToTwenty = len(password) >= 6 && len(password) <= 20

	return sixToTwenty && number && upper && lower
}
