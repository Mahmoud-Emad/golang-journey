package Maps
import "errors"



// We can always implement our custom type.


// type Dictionary map[string]string
// func (d Dictionary) Search (key string) (string, error) {
	// 	// This function should return the value of the key if it exists in the dictionary.
	// 	// return d[key], nil
	
	// 	// Now we have to check if the key is here 
	// 	definition, ok := d[key]
	// 	if !ok{
		// 		return "", errors.New("could not find the word you were looking for")
		// 	}
		// 	return definition, nil
		// }
		
		
		// Refactor
		
type Dictionary map[string]string

func NotFoundError() error {
	return errors.New("could not find the word you were looking for")
}
func (d Dictionary) Search (key string) (string, error) {
	definition, ok := d[key]
	if !ok{
		return "", NotFoundError()
	}
	return definition, nil
}