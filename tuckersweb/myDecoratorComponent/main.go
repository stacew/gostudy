package main

import (
	"fmt"
	"stacew/gostudy/tuckersweb/decoratorComponent2/cipher"
	"stacew/gostudy/tuckersweb/decoratorComponent2/lzw"
)

//Decorator interface
type Component interface {
	Operator(string) string
}

//1
type SendComponent struct {
	zipCom Component
}

func (self *SendComponent) Operator(data string) string {
	if self.zipCom != nil {
		data = self.zipCom.Operator(data)
	}

	return data
}

//2
type ZipComponent struct {
	encryptCom Component
}

func (self *ZipComponent) Operator(data string) string {
	//component -> zip
	if self.encryptCom != nil {
		data = self.encryptCom.Operator(data)
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
	unzipCom Component
}

func (self *ReadComponent) Operator(data string) string {
	if self.unzipCom != nil {
		data = self.unzipCom.Operator(data)
	}

	return data
}

//2
type UnzipComponent struct {
	decryptCom Component
}

func (self *UnzipComponent) Operator(data string) string {
	//unzip -> component
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}

	strUnzipData := string(unzipData)
	if self.decryptCom != nil {
		strUnzipData = self.decryptCom.Operator(strUnzipData)
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
		zipCom: &ZipComponent{
			//encryptCom : &EncryptComponent{
			//	key: "abcd",
			//},
		},
	}

	sendData := sender.Operator("Hello World")
	fmt.Println(sendData)

	receiver := &ReadComponent{
		unzipCom: &UnzipComponent{
			//decryptCom: &DecryptComponent{
			//	key: "abcd",
			//},
		},
	}

	recvData := receiver.Operator(sendData)
	fmt.Println(recvData)
}
