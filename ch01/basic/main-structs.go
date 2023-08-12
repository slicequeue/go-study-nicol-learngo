package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

// go 에서는 생성자가 없음! 스스로 실행해야함
// 언제? 이제 연습!~ Go에서는 struct 많이 알아야함! 메소드건 생성자건 다 struct 에서 부터 옴 - 그래서 잘 알아야!
// 이제 연습을 통해 알아보자!


func main() {
	nico := person{"nico", 18, []string{"kimchi", "ramen"}}
	slicequeue := person{name: "slicequeue", age: 33, favFood: []string{"kimchi", "ramen"}}
	fmt.Println(nico)
	fmt.Println(slicequeue)
}