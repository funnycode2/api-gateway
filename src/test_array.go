package main

func main() {
	/*a := make([]*A, 0)
	a = append(a, &A{})
	print(len(a))*/
	b := make([]I, 0)
	b = append(b, &A{})
	println(b[0].name())
}

type I interface {
	name() string
}

type A struct {
}

func (a *A) name() string {
	return "hyj"
}