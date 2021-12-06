/*
NOTE:
1. You have vendor in vendor/ but it only provide struct (we need interface for mocking),
so you need to create the interface manually / generator.

2. The interface already created in person/person.go,
run the go:generate to generate the mocks.
*/

package main

import (
	"fmt"

	person "github.com/verlandz/go-unitest/example/mock-vendor/person"
	vendor "github.com/verlandz/go-unitest/example/mock-vendor/vendor"
)

func Do(client person.Client, name string) string {
	client.SetName(name)
	return client.GetName()
}

// below is example to show relation between vendor(struct) and person(interface).
func main() {
	v := vendor.Person{}
	fmt.Println(Do(v, "Hello World"))
}
