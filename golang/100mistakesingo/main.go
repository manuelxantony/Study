package main

func main() {
	// embStruct()
	// embdInterfaceInStruct()
	// embeddSS2()
	// main1()
	// mainperson1()
	// mainSort()
	// embdInterfaceInStruct()
	// mainGetterSetter()
	// usingAny()
	// cmain()
	// ummain()

	// var counter uint = math.MaxUint
	// fmt.Println(counter + 1)

	// ----------------------------------------------------------------

	// s := make([]int, 3, 6)
	// fmt.Println(s)
	// s[0] = 1
	// s[1] = 2
	// s[2] = 3
	// fmt.Println(s)
	// // s[3] = 4  // -> error
	// s = append(s, 4)
	// fmt.Printf("%d %d addr: %p first: %p\n", s, cap(s), &s, &s[0])
	// s = append(s, 5)
	// s = append(s, 6)
	// fmt.Printf("%d %d addr: %p first: %p\n", s, cap(s), &s, &s[0])
	// s = append(s, 7)
	// fmt.Printf("%d %d addr: %p first: %p\n", s, cap(s), &s, &s[0])

	// -----------------------------------------------

	// s := make([]int, 5, 7)
	// s2 := s[1:3]
	// s[0] = 4
	// s[1] = 4
	// s[2] = 4
	// s[3] = 4
	// s[4] = 4
	// fmt.Println("s ", s, cap(s))
	// fmt.Println("s2", s2, cap(s2))
	// // s2[0] = 0
	// // fmt.Println(s)
	// s2 = append(s2, 5)
	// s2 = append(s2, 5)
	// s2 = append(s2, 5)
	// s2 = append(s2, 5)
	// fmt.Println("s ", s, cap(s))
	// fmt.Println("s2", s2, cap(s2))
	// s2[4] = 0
	// fmt.Println("s2", s2, cap(s2))
	// fmt.Println("s ", s, cap(s))

	// ------------------------------------------

	// s := make([]int, 3, 3)
	// s2 := s[:2]
	// s[0] = 1
	// s[1] = 1
	// s[2] = 1

	// fmt.Println("s", s, &s[0])
	// fmt.Println("s2", s2, &s2[0])
	// fmt.Println()

	// s2[0] = 5
	// fmt.Println("s", s, &s[0])
	// fmt.Println("s2", s2, &s2[0])
	// fmt.Println()

	// s = append(s, 1)
	// fmt.Println("s", s, &s[0])
	// fmt.Println("s2", s2, &s2[0])
	// fmt.Println()

	// s2[0] = 15
	// fmt.Println("s", s, &s[0])
	// fmt.Println("s2", s2, &s2[0])

	// ------------------------------------------

	// empty slice
	// s := make([]int, 10)
	// fmt.Println(len(s)) // 10
	// fmt.Println(cap(s)) // 10
	// fmt.Println(s[0])   // 0

	// // or
	// s2 := []int{}
	// fmt.Println(s2 == nil) // false
	// // fmt.Printf("%p", s2[0]) // error
	// fmt.Println("s2 cap", cap(s2)) // 0
	// s2 = append(s2, 1)
	// fmt.Println("s2 cap", cap(s2)) // 1
	// s2 = append(s2, 1)
	// fmt.Println("s2 cap", cap(s2)) // 2 doubled
	// s2 = append(s2, 1)
	// fmt.Println("s2 cap", cap(s2)) // 4 doubled
	// s2 = append(s2, 1)
	// fmt.Println("s2 cap", cap(s2)) // 4 len is samller = 3
	// s2 = append(s2, 1)
	// fmt.Println("s2 cap", cap(s2)) // 8 doubled

	// // nil slice
	// var s3 []int
	// fmt.Println(s3 == nil) // true
	// // fmt.Printf("%p", s3[0]) // error
	// fmt.Println(cap(s3)) // 0
	// fmt.Println(len(s3)) // 0

	// // or
	// s4 := []int(nil)
	// fmt.Println(s4 == nil) // true
	// // fmt.Printf("%p", s3[0]) // error
	// fmt.Println(cap(s4)) // 0
	// fmt.Println(len(s4)) // 0

	// s5 := append([]int(nil), 23)
	// fmt.Println(s5)

	// ---------------------------------------------------

	// copy
	// s := []int{1, 2, 3, 4, 5}
	// s2 := make([]int, 3)
	// copy(s2, s)
	// s2[0] = 0
	// fmt.Println(s2)
	// fmt.Println(s)

	// s := []int{1, 2, 2, 2, 2, 2, 2, 2, 2}
	// fmt.Println(len(s), cap(s))
	// s2 := s[:2]
	// fmt.Println(len(s2), cap(s2))
	// // s2[4] = 0 // err

	// s3 := make([]int, 4)
	// copy(s3, s)
	// fmt.Println(s3)
	// fmt.Println("s3 cap", cap(s3))
	// fmt.Println("s  cap", cap(s))

	// printMem()

	// s2 := s[:3]
	// fmt.Println(s2)
	// printMem()

	//runtime.GC()
	// memorymain()
	// runtime.GC()
	// fmt.Println()
	// memorymain2()
	// mainmap()
	// maincom()
	// mainCustomer()
	// mainLoopLabel()

	// i, j := returing()
	// fmt.Println(i, j)

	mainpanicrecover()

}

// func isNaN(f any) bool {
// 	return f != f
// }

// func returing() (i, j int) {
// 	j = 10
// 	return
// }
