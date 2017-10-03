package cli

import (
	"flag"
	"fmt"
	"os"
	"errors"
)

var nf = flag.Bool("newFamily", false, "Add a new family to database")
var na = flag.Bool("newAdult", false, "add an adult to an existing family")
var nc = flag.Bool("newChild", false, "add a child to an existing family")
var vd = flag.Bool("viewDatabase", false, "view the entire database")
var vf = flag.Bool("viewFamily", false, "view a families details")
var vm = flag.Bool("viewMember", false, "view a specific member")
var uf = flag.Bool("updateFamily", false, "update a family's information")
var um = flag.Bool("updateMember", false, "update a member's information")
var df = flag.Bool("deleteFamily", false, "remove a family from the database")
var da = flag.Bool("deleteAdult", false, "remove an adult form the database")
var dc = flag.Bool("deleteChild", false, "remove a child from the database")

func InitCLI() (string, []string, error)  {

	if len(os.Args) < 2 {
		return "", os.Args[1:], errors.New("please indicate an action e.g., -viewDatabase")
	}

	args := os.Args[2:]



	switch flag.Parse(); {
	case *nf: return "newFamily", args, nil
	case *na: return "newAdult", args, nil
	case *nc: fmt.Println("add new family")
	case *vd: fmt.Println("add new family")
	case *vf: fmt.Println("add new family")
	case *vm: fmt.Println("add new family")
	case *uf: fmt.Println("add new family")
	case *um: fmt.Println("add new family")
	case *df: fmt.Println("add new family")
	case *da: fmt.Println("add new family")
	case *dc: fmt.Println("add new family")
	default:
		return "", os.Args[1:], errors.New("action not recognized")
	}
	return "", os.Args[1:], errors.New("please indicate an action e.g., -viewDatabase")
}