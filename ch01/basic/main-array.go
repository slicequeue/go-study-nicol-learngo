package main

import "fmt"

func main3() {
	names := [5]string{"nico", "lynn", "dal"} //  배열은 크기 제한

	names[3] = "alala"
	names[4] = "alala"
	// names[5] = "alala" // error

	fmt.Println(names)

	sliceNames := []string{"nico", "lynn", "dal"} // 슬라이스 타입 크기를 지정하지 않음
	// sliceNames[3] = "lalala" // error 이건 안됨 
	sliceNames = append(sliceNames, "lalal") // 값은 불변하게 새로운 슬라이스를 리턴함!

	fmt.Println(sliceNames)

}