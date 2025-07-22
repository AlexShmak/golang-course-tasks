package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) getName() string {
	return h.name
}

func (h *Human) getAge() int {
	return h.age
}

type Action struct {
	Human
	title string
}

func (a *Action) getTitle() string {
	return a.title
}

func main() {
	action := Action{
		Human: Human{name: "TestName", age: 21},
		title: "TestActionTitle",
	}
	fmt.Println(action.getName())
	fmt.Println(action.getAge())
	fmt.Println(action.getTitle())
}
