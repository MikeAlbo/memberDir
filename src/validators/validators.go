//Validators package provides a set of validators for length, character, and int
package validators

import (
	"errors"
	"regexp"
)

var (
	intRegex = regexp.MustCompile(`^[0-9]*$`)
	charRegex = regexp.MustCompile(`^[a-zA-Z]*$`)
)

// length validator
func StringLengthValidator(s string, length int)(string, error){
	if len(s) != length {
		return "", errors.New("input is of invalid length")
	}
	return s, nil
}

// int validator
func IntValidator(val string)(string, error)  {
	if intRegex.MatchString(val) {
		return val, nil
	}

	return "", errors.New("value should be an integer")
}

// char validator
func CharValidator(val string)(string, error)  {
	if charRegex.MatchString(val) {
		return val, nil
	}

	return "", errors.New("value should be a non-integer character")
}
