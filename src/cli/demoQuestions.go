package cli

type Question struct{
	prompt string
	minLength int
	maxLength int
	inputType string
}


var Name = &Question{
	prompt: "What is your Last Name?",
	minLength: 1,
	maxLength: 15,
	inputType: "string",
}

var Address  = &Question{
	prompt: "What is your street address?",
	minLength: 3,
	maxLength: 30,
	inputType: "address",
}

var Apartment = &Question{
	prompt: "apartment number? What is your apartment number?",
	minLength: 1,
	maxLength: 5,
	inputType: "apartment",
}

var City = &Question{
	prompt:"What City?",
	minLength: 3,
	maxLength: 13,
	inputType: "string",
}

var State = &Question{
	prompt:"What State (TN, MI, CA, ...)?",
	minLength: 2,
	maxLength: 2,
	inputType:"string",
}

var zipcode = &Question{
	prompt: "What is your zipcode?",
	minLength: 5,
	maxLength: 5,
	inputType: "int",
}

func getNewFamilyQuestions() []Question  {
	return []Question{*Name, *Address, *Apartment, *City, *State, *zipcode }
}