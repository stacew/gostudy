package main

import (
	"fmt"

	"gostudy/nomadcoders/2.1.DictionaryProjects/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	err := dictionary.Add(word, "Hello World")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)

	definition, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add(word, "New Hello World")
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Update(word, "New Hello World")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Update Success")
	}

	definition, err = dictionary.Search(word)
	fmt.Println(definition)

	err = dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	}
	err = dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	}

	definition, err = dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
