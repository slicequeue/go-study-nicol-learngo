package main_hold

import (
	"fmt"
	"time"
)

func main_hold() {
	// 메인함수는 다른 goroutines 를 기다려주지 않는다
	// 메인함수가 끝나면 goroutines 들도 다 소멸됨
	// 어떻게 메인 함수에서 고루틴들이 커뮤니케이션 할지 그것을 우리는 아직 모른다... 
	// 채널을 통해서 하는 것! 메인함수는 결과를 저장하는 곳! 
	// 이를 위해서는 메인 함수와 고루틴간 커뮤니케이션 

	// 채널은 메인함수와 고루틴간의 커뮤니케이션, 고루틴과 고루틴간에 커뮤니케이션 가능!
	// 어떻게 커뮤니케이션 하는가?
	// 채널은 파이프 같은 것
	people := [2]string{"nico", "flynn"}
	c := make(chan bool) // 원하는 채널 그리고 주고 받는 데이터 형식 만들기
	for _, person := range people {
		// result := go isSexy(person) // 이렇게 할 수 없다
		go isSexy(person, c)
	}
	fmt.Println(<- c) // go 는 똑똑해서 이미 두개 고루틴이 도는 것을 알고 있어서 이렇게 하면 에러가 발행함 <- c 세번 했으니!...
	fmt.Println(<- c) // 이처럼 두번 하면 됨
	fmt.Println(<- c) // 하나만 결과가 나옴! 두개를 받을 수 있을까?
	// result := <- c // 결과를 꺼내서 result 로 ! 초기화
	// fmt.Println(result) // 아래 10초 슬립 없이도 결과를 기다림!
	// time.Sleep(time.Second * 10)


}

func isSexy(person string, c chan bool) /* bool 어짜피 안됨*/ {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true // 화살표로 보낸다 표시!
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}