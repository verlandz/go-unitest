package mock_struct

import (
	person "github.com/verlandz/go-unitest/unitest-sample/mock-struct/person"
)

// ASSUMING YOUR DIRECTORY IS HERE.

// 1. Assuming you have vendor that only have struct and func
// without interface.
// 	ex: mock-struct/vendor

// 2. You want to use the vendor and do mock unitest it.
// since there's no interface provided,
// you should generate the interface from it.
// 	cmd: mkdir person && ifacemaker -f vendor/vendor.go -s Person -i Person -p person -o person/Person.go
//
// actually you can create by yourself without using generator.
// if possible, add also cmd in comment for generate mockgen
// 	//go:generate mockgen -source=Person.go -destination=mocks/Person.go -package=mocks . Person

// 3. Mock the interface
//	execute "//go:generate mockgen ..." from person/Person.go

// 4. From this following usecase, try to mock it
func Do(p person.Person, name string) string {
	p.SetName(name)
	return p.GetName()
}
