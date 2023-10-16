package main

import (
	"fmt"
)

type PersonInterface interface {
	addName(string)
}

type Student1 struct {
	name string
}

type Teacher1 struct {
	name string
}

func (s *Student1) addName(name string) {
	s.name = name
}

func (t *Teacher1) addName(name string) {
	t.name = name
}

func addName(p PersonInterface, name string) {
	p.addName(name)
}

func mainperson1() {
	s := &Student1{}
	addName(s, "biju")

	t := &Teacher1{}
	addName(t, "mon")

	fmt.Println(s.name, t.name)

}
