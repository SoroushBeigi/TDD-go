package dictionary

import "errors"

var errNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errNotFound
	} else {
		return definition, nil
	}
}

func (d Dictionary)Add(word,definition string){
	
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
