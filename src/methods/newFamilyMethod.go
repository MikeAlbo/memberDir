// New Family Method
package methods

import (
	"cli"
	"fmt"
)

func AddNewFamily(family string)  {
	fmt.Printf("Adding the %s family to the database\n", family)
	if f, err := cli.PromptUser("newFamily"); err != nil {
		panic(err)
	} else {
		fmt.Println(f)
	}
}


// un-marshall JSON into data, validates that family doesn't already exist
// if exist, concats first name and last name into family field i.e., Smith Smith,Mark, Smith,Tony
// create a new "family" from the family struct

// -- below should handle transactions with the CLI --
// prompt user that a new family (name) is being added to the db.
// call user interface with [] of questions for the prompts

// -- handles the formatting of the data and writing to the db --
// takes an []args and builds out the family struct
// prompts user to confirm (method) yes, no, ""
// calls db functions to marshall struct > JSON
// writes the JSON to the file
//