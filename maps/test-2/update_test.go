package Maps
import "testing"



// Next will be the update method for testing
func TestUpdate(t *testing.T){
	dictionary := Dictionary{}
	key := "test"
	value := "This is a value test"
	dictionary.Add(key, value)
	newValue := "This is a new value test"
	dictionary.Update(key, newValue)
	assertDefinition(t, dictionary, key, newValue)
}

// There is no refactoring we need to do on this since it was a simple change. However, we now have the same issue as with Add. 
// If we pass in a new word, Update will add it to the dictionary.

func TestUpdateV2(t *testing.T) {
	key := "test"
	value := "This is a value test"
	dictionary := Dictionary{key: value}
	
	t.Run("Update existing key", func(t *testing.T) {
		newValue := "This is a new value test"
		err := dictionary.Update(key, newValue)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newValue)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
	
		err := dictionary.Update(word, definition)
	
		assertError(t, err, ErrWordDoesNotExist)
	})
}