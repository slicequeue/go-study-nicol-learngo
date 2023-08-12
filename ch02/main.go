package main

import (
	"fmt"
	"log"

	"github.com/slicequeue/learngo/ch02/accounts"
	"github.com/slicequeue/learngo/ch02/mydict"
)

func main_hold() {
	account := accounts.NewAccount("nico")
	// account.balance = 100 // 이게 안됨!!
	fmt.Println(*account)
	account.Deposit(1000)
	fmt.Println(account.Balance())
	// 이게 고에서 처리하는 에러 방식! 항상 에러를 체크해야함 - error 체크를 강제! 어찌보면 손수 처리하는데 이상한데... 이게 좋다는데?!?
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err) // Println 후 종료시켜줌
	}
	fmt.Println(account.Balance())
	account.Withdraw(200)

	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	word2 := "hello"
	definition2 := "Greeting"
	err2 := dictionary.Add(word2, definition2)
	if err2 != nil {
		fmt.Println(err2)
	}
	hello, _ := dictionary.Search(word2)
	fmt.Println("fount:", word2, "definition:", hello)
	err3 := dictionary.Add(word2, definition2)
	if err3 != nil {
		fmt.Println(err3)
	}

	err4 := dictionary.Update(word2, "Second")
	if err4 != nil {
		fmt.Println(err4)
	}
	definition3, _ := dictionary.Search(word2)
	fmt.Println(definition3)

	err5 := dictionary.Delete(word2)
	if err5 != nil {
		fmt.Println(err5)
	}
	_, err6 := dictionary.Search(word2)
	if err6 != nil {
		fmt.Println(err6)
	}
}