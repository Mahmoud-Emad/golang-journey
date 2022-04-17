package Maps


// Write the type first
type Dictionary map[string]string
type DictionaryError string


const (
	ErrNotFound         = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryError("cannot update word because it does not exist")
)

func (err DictionaryError) Error() string {
	return string(err)
}

func (d Dictionary)Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary)Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok{
		return "", ErrNotFound
	}
	return definition, nil	
}

func (d Dictionary) Delete(key string) (string, error) {
	definition, err := d.Search(key)
	if err != nil {
		return "", err
	}
	delete(d, key)
	return definition, nil
}