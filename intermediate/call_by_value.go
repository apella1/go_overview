package intermediate

type person struct {
	name string
	age  int
}

func ModifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}
