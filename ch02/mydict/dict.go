package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound error = errors.New("Not Found")
var errWordExists = errors.New("That word ")
var errCantUpdate = errors.New("Can't update non-exisiting word")
var errCantDelete = errors.New("Can't delete non-exisiting word")

// struct 에만 하는 게 아니라 타입에도 가능한 리시버 

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	// return "", errors.New("Not Found")
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
