package cat

type Cat struct {
	Name string
	Years int
}

func (c Cat) Age() int {
	return c.Years
}
