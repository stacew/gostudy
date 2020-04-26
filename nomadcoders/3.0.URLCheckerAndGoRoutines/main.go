package main

import (
	"errors"
	"fmt"
	"net/http"
)

type reqResult struct {
	url    string
	status string
}

var errRequestFail = errors.New("Request Failed")

func main() {

	var results = make(map[string]string)
	c := make(chan reqResult)

	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.daum.com",
		"https://www.instagram.com",
		"https://www.facebook.com",
		"https://www.reddit.com",
		"https://www.stacew.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for range urls {
		recvURL(c)
	}

}

//chan<- Input Only Type
func hitURL(url string, c chan<- reqResult) {
	//fmt.Println("Checking:", url)

	status := "OK"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- reqResult{url: url, status: status}
}

func recvURL(c <-chan reqResult) {
	result := <-c
	fmt.Println(result.url, result.status)
}
