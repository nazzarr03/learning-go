package email_validator

import (
	"regexp"	
)

var emailRegexp = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {	
	if len(email) > 254 {
		return false
	}
	return emailRegexp.MatchString(email)
}