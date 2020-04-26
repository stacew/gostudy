package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That Word Exists")

//Search for word
//slice, channel, map은 reciver 붙일 때 기본이 refrenceType이라서 d Dictionry
func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if ok {
		return value, nil
	}

	return "", errNotFound
}

//Add ~~
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExists
		}
	*/

	return nil
}

//Update ~~
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNotFound
	}

	d[word] = def

	return nil
}

//Delete ~~
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNotFound
	}

	delete(d, word)
	return nil
}
