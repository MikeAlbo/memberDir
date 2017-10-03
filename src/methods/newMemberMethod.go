// Add a new member to an existing family

package methods

import "fmt"

func AddNewAdult(args []string)  {
	fmt.Printf("Adding %s, to the %s family", args[1], args[0])
}
