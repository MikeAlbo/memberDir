package main

import (
	"encoding/json"
	"fmt"
	"dbService"
	"memberModel"
	"cli"
	"os"
	"methods"
)



func main()  {

	action, args, err := cli.InitCLI()
	softExit(err, "For additional options, use -help")
	chooseMethod(action,args)


	michael := &memberModel.Member{
		FirstName: "Michael",
		LastName: "Albonetti",
		Email: "malbo24@gmail.com",
		CellPhone: "901-2181977",
		BirthDate: "06-25-1984",
	}

	kristy := &memberModel.Member{
		FirstName: "Kristy",
		LastName: "Horton",
		Email: "klh@yahoo.com",
		CellPhone: "6165761243",
		BirthDate: "06-10-1983",
	}

	moose := &memberModel.Member{
		FirstName: "Moose",
		LastName: "Money",
		Email: "dog1@yahoo.com",
		CellPhone: "",
		BirthDate: "04-15-2015",
	}

	patchy := &memberModel.Member{
		FirstName: "Patchy",
		LastName: "Patch",
		Email: "spots@yahoo.com",
		CellPhone: "",
		BirthDate: "06-15-2014",
	}

	testFamily := &memberModel.Family{
		Address: "614 blue lake trail",
		City: "Antioch",
		State: "TN",
		Zip: 37013,
		Adults: []memberModel.Member {*michael, *kristy},
		Children: []memberModel.Member {*moose, *patchy},
	}

	var testFamilies []interface{}

	for i:=0; i < 5; i++ {
		testFamilies = append(testFamilies, *testFamily)
	}

	testFamRes, err := json.MarshalIndent(testFamilies, "", " ")
	errCheck(err)

	if n,err := dbService.SetData(testFamRes); err != nil {
		fmt.Errorf("error: %v\n", err)
	} else {
		fmt.Printf("data was written to %s successfully\n", n)
	}

}

func errCheck(e error)  {
	if e != nil {
		panic(e)
	}
}

func softExit(e error, m string)  {
	if e != nil {
		fmt.Printf("%s\n%s\n",e,m)
		os.Exit(1)
	}
}

func chooseMethod(action string, args []string) {
	switch action {
	case "newFamily": methods.AddNewFamily(args[0])
	case "newAdult": methods.AddNewAdult(args[0:])

	}
}
