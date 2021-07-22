package vendor

type Person struct {
	name string
}

// SetName set person's name.
func (p Person) SetName(name string) {
	p.name = name
}

// GetName get person's name.
func (p Person) GetName() string {
	return p.name
}
