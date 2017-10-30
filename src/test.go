package main

import (
	"reflect"
)

func main() {
	benz := &Benz{
		model: "S600L AMG",
	}
	elem := reflect.ValueOf(&benz).Elem()
	model := elem.MethodByName("Buy").Call(nil)[0]
	println("i want to buy " + model.String() + ", but")
	funcChangeIdea := elem.MethodByName("ChangeIdea")
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("v40")
	funcChangeIdea.Call(params)
	models := elem.MethodByName("Buy").Call(nil)[0]
	println("i buy " + models.String() + " instead")
}

type Benz struct {
	model string
}

func (b *Benz) Buy() string {
	return b.model
}

func (b *Benz) ChangeIdea(idea string)  {
	b.model = idea
}