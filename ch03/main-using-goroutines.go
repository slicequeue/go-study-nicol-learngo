package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

func main() {
	var results = make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i:=0; i< len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("Request failed")

func hitURL(url string, c chan<- requestResult /* `chan<-` 이 채널을 보낼수 만 있고 받을 수 없다 표시 */) error {
	// fmt.Println(<- c) // 이렇게 받을 수 있지만 `chan<-` 으로 인해 에러
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} 
	c <- requestResult{url:url, status: status}
	return nil
}
