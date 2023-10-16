package main

import (
	"fmt"
)

func (g *Graph) DFS(source string) {
	var s Stack

	s.Push(source) // pushed a

	for !s.IsEmpty() {
		poped := s.Pop()
		fmt.Println(poped)
		neigh := g.Map[poped]
		s.Push(neigh...)
	}
}

func (g Graph) DFSRecur(source string) {
	fmt.Println(source)
	for _, neig := range g.Map[source] {
		g.DFSRecur(neig)
	}

}
