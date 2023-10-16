package main

import "fmt"

// parent
type GetterSetterInterface interface {
	Get() int // or GetterInterface
	Set() int
}

type GetSet struct {
	value int
}

func (g *GetSet) Get() int {
	return g.value
}

func (g *GetSet) Set(value int) {
	g.value = value
}

// now we are having an implementation which
// we need to restrict the Set

// child, because of the signature
type GetterInterface interface {
	Get() int
}

type Getter struct {
	getter GetterInterface
}

func mainGetterSetter() {
	g := &GetSet{}
	// in parent set is exposed
	g.Set(10)

	// in the following we can only access Get
	gNew := &Getter{getter: g}
	fmt.Println(gNew.getter.Get())
}
