package structs

// Pet type
type Pet struct {
	Name string
	Kind string
}

// MyPet function
func MyPet() Pet {
	var p Pet
	p.Name = "Frankies"
	p.Kind = "dog"
	return p
}
