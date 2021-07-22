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
// 	cmd: mkdir person && ifacemaker -f vendor/vendor.go -s Person -i Person -p person -o person/IPerson.go

// 3. Mock the interface
//	cmd: mockgen -source=person/IPerson.go -destination=mocks/mock.go -package=mocks

// 4. From this following usecase, try to mock it
func Do(p person.Person, name string) string {
	p.SetName(name)
	return p.GetName()
}
