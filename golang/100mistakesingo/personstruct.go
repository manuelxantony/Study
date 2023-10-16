package main

import "fmt"

type Person struct {
	name string
}

type Student struct {
	Person
}

type Teacher struct {
	Person
}

func (s *Person) addName(name string) {
	s.name = name
}

func mainperson() {
	s := Student{}
	s.addName("vava")

	t := Teacher{}
	t.addName("suresh")

	fmt.Println(s)
	fmt.Println(t)
}
