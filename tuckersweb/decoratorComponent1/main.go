package main

import (
	"fmt"
	"stacew/gostudy/tuckersweb/decoratorComponent1/cipher"
	"stacew/gostudy/tuckersweb/decoratorComponent1/lzw"
)

//interface
type Component interface {
	Operator(string) string
}

//1
type SendComponent struct{}

func (self *SendComponent) Operator(data string) string {
	// Send data
	return data
}

//2
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) string {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	return self.com.Operator(string(zipData))
}

//3
type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) string {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	return self.com.Operator(string(encryptData))
}

/////////////////////////////////////////////////////
//1
type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) string {
	return data
}

//2
type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) string {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	return self.com.Operator(string(unzipData))
}

//3
type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) string {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	return self.com.Operator(string(decryptData))
}

//main
func main() {
	sender :=
		//&EncryptComponent{
		//key: "abcd",
		&ZipComponent{
			com: &SendComponent{},
			//},
		}

	sentData := sender.Operator("Hello World")
	fmt.Println(sentData)

	receiver :=
		&UnzipComponent{
			//com: &DecryptComponent{
			//	key: "abcd",
			com: &ReadComponent{},
			//},
		}

	recvData := receiver.Operator(sentData)
	fmt.Println(recvData)
}
