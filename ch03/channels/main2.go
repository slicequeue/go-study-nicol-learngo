package main

import (
	"fmt"
	"time"
)

func main_hold() {
	// 채널은 메인함수와 고루틴간의 커뮤니케이션, 고루틴과 고루틴간에 커뮤니케이션 가능!
	// 어떻게 커뮤니케이션 하는가?
	// 채널은 파이프 같은 것
	people := []string{"nico", "flynn", "dal", "japanguy", "larry"}
	c := make(chan string) // 원하는 채널 그리고 주고 받는 데이터 형식 만들기
	for _, person := range people {
		// result := go isSexy(person) // 이렇게 할 수 없다
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ { // 이렇게 고루틴 갯수를 알 수 있는 것을 통해 포문 처리
		fmt.Print("waiting for ", i, " ")
		fmt.Println(<-c) 
	}

	// 룰을 기억하자! 채널의 룰은 매우 심플! 채널이랑 고루틴! 
	// 첫째 메인함수가 종료되면 고루틴은 의미가 없음 
	// 두번째 너가 받을 데이터에 대해서 어떤 타입을 받을지 구체적으로 지정해줘야함
	// 세번째 메세지를 채널로 보내는 방법은 화살표를 채널로 향해 처리 c <- data

	// 기존 방식 일일히 하나씩 받는 처리
	// resultOne := <- c    // blocking operation - 메세지 받으면 다음 라인으로 넘어감!
	// resultTwo := <- c  	 // 이 줄도 마찬가지! - 메세지 받을때 까지 대기 한다는 것!
	// resultThree := <- c  // 이 줄도 마찬가지! - 메세지 받을때 까지 대기 한다는 것!
	// fmt.Println("Waiting for messages")
	// fmt.Println("Received this message:", resultOne)
	// fmt.Println("Received this message:", resultThree)
	
}

func isSexy(person string, c chan string) /* bool 어짜피 안됨*/ {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"// 화살표로 보낸다 표시!
}
