package main

import (
	"fmt"
	"time"
)

func main() {
	go tenCount("nico")
	go tenCount("yslee")
	time.Sleep(time.Second * 3) //중요! main 함수는 goroutine을 기다려주지 않는다.
}

func tenCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, ":", i)
		time.Sleep(time.Second)
	}

}
