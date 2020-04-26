package main

func main() {
	ch := make(chan string, 1)
	sen1dChan(ch)
	rec1eiveChan(ch)
}

func sen1dChan(ch chan string) {
	ch <- "Data"
	x := <-ch // 엝러발생
	print(x)
}

func rec1eiveChan(ch chan string) {
	//data := <-ch
	//fmt.Println(data)
}
