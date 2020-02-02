package test

type test_s struct {
	A int
}

func NeWTest(a int) *test_s {
	return &test_s{a}
}

type person struct {
	fitstName, lastName string
}

func NewPerson(a, b string) *person {
	return &person{a, b}
}

func (p *person) GetFirstName() string {
	return p.fitstName
}
