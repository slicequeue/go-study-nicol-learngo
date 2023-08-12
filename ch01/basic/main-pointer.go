package main

import "fmt"

func main2() {
	// a := 2
	// b := a // copy value!!!
	// a = 10
	// fmt.Println(a, b)   // 10, 2
	// fmt.Println(&a, &b) // 0xc0000a6058 0xc0000a6070

	a := 2
	b := &a
	a = 5
	fmt.Println(&a, b) // 0xc000016098 0xc000016098
	fmt.Println(*b)    // 5
	*b = 20
	fmt.Println(a) // 20
}
