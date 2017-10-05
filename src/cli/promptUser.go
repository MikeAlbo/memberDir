// The prompt user section of the cli method controls the input flow of the user
package cli

import (
	"validators"
	"errors"
	"bufio"
	"os"
	"fmt"
)

// PromptUser takes in an action and returns a slice of strings (answers from the user),
// or an error that may have bubbled up from dependent methods
func PromptUser(action string) ([]string, error) {
	questions, err := selectQuestionSet(action)
	if err != nil {
		panic(err) // todo modify to global error handler 
	}
	
	var results []string
	
	for _,q := range questions {
		fmt.Println(q.prompt)
		ans, err := activateScanner()
		if err != nil {
			panic(err)
		}

		if vI, err := validateInput(ans, q); err !=nil {
			panic(err)
		} else {
			results = append(results, vI)
		}
	}

	return results,nil
}

// selectQuestionSet takes in an action and returns the set of questions related to the action
// will return an error if the action does not match a predefined case
func selectQuestionSet(action string) ([]Question, error)  {
	switch action {
	case "newFamily": return getNewFamilyQuestions(), nil
	default:
		return nil, errors.New("incorrect action given inside cli > promptUser > selectQuestion")
	}
}

// activateScanner init the bufio.Scanner and returns the user's input
// will pass an error up the stack if one exist inside the scanner
func activateScanner() (string, error)  {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		return scanner.Text(), nil
	}
	if  err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

// validateInput validates the input given by the user vs  what is required by the question
// will return an error if question.inputType does not match a case.
func validateInput(input string, q Question) (string,error) {

	var validInputType string
	var err error = nil

	switch q.inputType {
	case "string": validInputType, err = validators.CharValidator(input)
	case "int": validInputType, err = validators.IntValidator(input)
	case "address": validInputType, err = validators.AddressValidator(input)
	case "apartment": validInputType, err = validators.ApartmentValidator(input)
	default:
		return "",errors.New("bad validation type")
	}

	if err != nil {
		return "", err
	}

	if _, err := validators.StringLengthValidator(validInputType, q.minLength, q.maxLength); err != nil {
		return "", err
	}

	return validInputType, nil
}



