package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrorNotFound        = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists      = DictionaryErr("cannot add wordbecause it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	/*
		In order to make this pass, we are using an interesting property of the map lookup. It can return 2 values.
		The second value is a boolean which indicates if the key was found successfully.
	*/
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordExists
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrWordDoesNotExists
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
