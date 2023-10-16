package main

// import "fmt"

// type InterfaceA interface {
// 	Len() int
// }

// type InterfaceB interface {
// 	Size() int
// }

// // here only B satisfies Interface, as B satisfies both interfaceA and interfaceB
// type Interface interface {
// 	InterfaceA
// 	InterfaceB
// }

// // A explicitly satisfing interfaceA
// type A struct {
// 	len int
// }

// // B implicitly satisfies InterfaceA by ebedding A
// // and explicitly satisfing InterfaceB
// type B struct {
// 	A
// 	len  int
// 	size int
// }

// type C struct {
// 	A
// 	B
// }

// func (a A) Len() int {
// 	return a.len
// }

// func (b B) Len() int {
// 	return b.len
// }

// func (b B) Size() int {
// 	return b.size
// }

// func printSS(i InterfaceA) {
// 	fmt.Println("len:", i.Len(), i)
// }

// func printSSB(i InterfaceB) {
// 	fmt.Println("size", i.Size(), i)
// }

// func printCommonInterface(i Interface) {
// 	fmt.Println(i.Len())
// 	fmt.Println(i.Size())
// }

// func embeddSS2() {
// 	a := A{5}
// 	printSS(a)

// 	b := B{a, 10, 6}
// 	printSS(b)
// 	printSSB(b)

// 	printCommonInterface(b)

// 	c := C{a, b}
// 	printSS(c)
// }
