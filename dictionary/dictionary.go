package dictionary

import "errors"

var (
	errNotFound               = errors.New("could not find the word you were looking for")
	errWordExists             = errors.New("cannot add word because it already exists")
	errWordDoesNotExistUpdate = errors.New("cannot update word because it does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errNotFound
	} else {
		return definition, nil
	}
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errWordDoesNotExistUpdate
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
