package main

import (
	"fmt"
	"strings"
)

var name2 string = "nico"

func repeatMe(words ...string) { //... list 무제한 Input
	fmt.Println(words)
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("i'm Done!")
	fmt.Println("I'm Run!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func addUseFor(numbers ...int) int {
	for index := range numbers {
		fmt.Println(index)
	}
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	//if에서만 쓰는 변수 생성
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else { // else 필요 없다고 경고 뜸
		return true
	}
}

func canIDrink2(age int) bool {
	//go는 c++ switch-case처럼 점프하는 게 없는듯?
	//if elif else 등의 코드를 축약사용 위해 그러는듯...
	//Default로 Break가 숨겨진 느낌임.
	//그래서 case: 위 아래로 이어 붙이면 큰일남.
	//아래 case:로 확인하고 돌아간다 생각하면 real super bug

	switch test := 18; {
	case test > 18:
		return true
	default:
		//return false
	}

	test := 18
	switch { //지울수있음
	case test == 18: //바로 밑에 내용 없어서 무시되버림
	case test > 18:
		return true
	case test == 17:
	case test == 17:
	default:
		return false
	}

	switch koreanAge := age + 2; koreanAge { //좌 초기화, 우(마지막?) 조건
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	const cName = "notype"
	const cNameS string = "type"

	name := "nico" //축약형은 함수 내부만 가능
	fmt.Println(name)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	totalLength, uppercase := lenAndUpper("nico")
	fmt.Println(totalLength)
	fmt.Println(uppercase)

	fmt.Println("addUseFor", addUseFor(10, 20, 30))

	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink2(16))

	a := 2
	b := &a
	*b = 1000
	fmt.Println(a)

	names := []string{"nico", "lynn", "yslee"}
	names2 := append(names, "ggg") //리턴값 받는 처리 꼭.
	fmt.Println(names)
	fmt.Println(names2)

	//unordered
	maps := map[string]string{
		"ee": "ee",
		"dd": "dd",
		"cc": "cc",
		"bb": "bb",
		"aa": "11"}
	fmt.Println(maps) //정렬된 결과 출력
	//for unordered 출력
	for key, value := range maps {
		fmt.Println(key, value)
	}

	favFood := []string{"ki", "ra"}
	ys1 := Person{name: "yslee", age: 100, favFood: favFood}
	fmt.Println(ys1)
}

type Person struct {
	name    string
	age     int
	favFood []string
}
