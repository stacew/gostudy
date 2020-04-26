package main

import (
	"fmt"
	"gostudy/tucker/decoratorComponent2/cipher"
	"gostudy/tucker/decoratorComponent2/lzw"
)

//interface
type Component interface {
	Operator(string) string
}

//1
type SendComponent struct {
	com Component
}

func (self *SendComponent) Operator(data string) string {
	if self.com != nil {
		data = self.com.Operator(data)
	}

	return data
}

//2
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) string {
	//component -> zip
	if self.com != nil {
		data = self.com.Operator(data)
	}

	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	return string(zipData)
}

//3
type EncryptComponent struct {
	key string
}

func (self *EncryptComponent) Operator(data string) string {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	return string(encryptData)
}

/////////////////////////
//1
type ReadComponent struct {
	com Component
}

func (self *ReadComponent) Operator(data string) string {
	if self.com != nil {
		data = self.com.Operator(data)
	}

	return data
}

//2
type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) string {
	//unzip -> component
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}

	strUnzipData := string(unzipData)
	if self.com != nil {
		strUnzipData = self.com.Operator(strUnzipData)
	}

	return strUnzipData
}

//3
type DecryptComponent struct {
	key string
}

func (self *DecryptComponent) Operator(data string) string {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	return string(decryptData)
}

//main
func main() {
	sender := &SendComponent{
		com: &ZipComponent{
			//com: &EncryptComponent{
			//	key: "abcd",
			//},
		},
	}

	sendData := sender.Operator("Hello World")
	fmt.Println(sendData)

	receiver := &ReadComponent{
		com: &UnzipComponent{
			//com: &DecryptComponent{
			//	key: "abcd",
			//},
		},
	}

	recvData := receiver.Operator(sendData)
	fmt.Println(recvData)
}
