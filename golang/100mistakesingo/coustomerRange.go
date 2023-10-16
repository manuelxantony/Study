package main

import "fmt"

type Coustomer struct {
	ID   int
	Name string
}

type Store struct {
	m map[int]*Coustomer
}

// wrtong range implementation
func (s *Store) storeCoustomers(coustomer []Coustomer) {
	for _, c := range coustomer {
		// every time c is going to have same address
		// that is when ranging over coustomers the c is storing the value in same address
		s.m[c.ID] = &c
	}
}

// wrtong range implementation
func (s *Store) storeCoustomersRight(coustomer []Coustomer) {
	for _, c := range coustomer {
		// coping the value in the c to current
		current := c
		s.m[c.ID] = &current
	}
}

func mainCustomer() {
	c := []Coustomer{
		{1, "bob"},
		{2, "chandu"},
		{3, "shahsi"},
	}
	m := make(map[int]*Coustomer)
	s := Store{m}
	s.storeCoustomers(c)
	fmt.Println(s)

	m = make(map[int]*Coustomer)
	s = Store{m}
	s.storeCoustomersRight(c)
	fmt.Println(s)
}
