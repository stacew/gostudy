package main

import "time"

func main() {
	c := make(chan string)
	people := []string{"nico", "yslee", "yslee2", "yslee3", "yslee4", "yslee5", "yslee6", "yslee7"}
	for _, person := range people {
		go isWait(person, c)
	}

	for i := range people {
		println(i, <-c)
	}

	//println(<-c) //deadlock
}

func isWait(person string, c chan string) {
	time.Sleep(time.Second * 5)
	//5초 뒤에 보내도 goroutines 처리 가능
	c <- person + " : End"
}
