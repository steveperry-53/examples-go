package dog

type Dog struct {
	Name string
	Pounds int
}

func (d Dog) Weight () int {
	return d.Pounds
}
