package main

import (
	"fmt"
	"time"
)

func main_hold() {
	people := [2]string{"nico", "flynn"}

	go sexyCount("nico")
	go sexyCount("flynn")

	time.Sleep(time.Second * 5)

	// 메인함수는 다른 goroutines 를 기다려주지 않는다
	// 메인함수가 끝나면 goroutines 들도 다 소멸됨
	// 어떻게 메인 함수에서 고루틴들이 커뮤니케이션 할지 그것을 우리는 아직 모른다... 
	// 채널을 통해서 하는 것! 메인함수는 결과를 저장하는 곳! 
	// 이를 위해서는 메인 함수와 고루틴간 커뮤니케이션 

	// 채널은 메인함수와 고루틴간의 커뮤니케이션, 고루틴과 고루틴간에 커뮤니케이션 가능!
	// 어떻게 커뮤니케이션 하는가?
	// 채널은 파이프 같은 것

}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}