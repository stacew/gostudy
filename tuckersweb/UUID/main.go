package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New() //유니버셜 유니크 아이디 : 우주 상에 동일한 키가 나올 수 없다..?
	fmt.Println(id.String())
}
