//Validators package provides a set of validators for length, character, and int
package validators

import (
	"errors"
	"regexp"
)

var (
	intRegex = regexp.MustCompile(`^[0-9]*$`)
	charRegex = regexp.MustCompile(`^[a-zA-Z]*$`)
	apartmentRegex = regexp.MustCompile(`([A-Za-z0-9])\w+`)
	addressRegex = regexp.MustCompile(`^\d+\s[A-z]+\s[A-z]+`)
)

// length validator
func StringLengthValidator(s string, minLength, maxLength int)(string, error){
	if len(s) < minLength || len(s) > maxLength {
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

func AddressValidator(val string)(string, error)  {
	if addressRegex.MatchString(val){
		return val, nil
	}

	return "", errors.New("value should be a properly formatted address")
}

func ApartmentValidator(val string)(string, error) {
	if apartmentRegex.MatchString(val){
		return val, nil
	}

	return "", errors.New("apartment numbers should only contain letters and numbers")
}
