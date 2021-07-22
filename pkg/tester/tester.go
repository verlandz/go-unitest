package tester

import (
	"fmt"
)

type Name struct {
	Given, When, Then string
}

func (n Name) Construct() string {
	if n.Given == "" || n.When == "" || n.Then == "" {
		panic("(Given|When|Then) can't be empty.")
	}

	return fmt.Sprintf("Given %v. When %v. Then %v.", n.Given, n.When, n.Then)
}
