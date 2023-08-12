package main2

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url string
	status int
}

func main_hold() {
	// var results map[string]string // 초기화 하지 않은 맵에 값을 할당하면 오류
	// results["gello"] = "Hello"
	var results = make(map[string]string)
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("Request failed")

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	// go lang std library
	// 대부분 https://pkg.go.dev/std 다 있음
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
