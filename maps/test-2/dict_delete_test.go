package Maps
import "testing"

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test":"this is just a test"}
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		_, err := dictionary.Delete(word)
		assertError(t, err, nil)
		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("delete non-existing word", func(t *testing.T) {
		_, err := dictionary.Delete("test2")
		assertError(t, err, ErrNotFound)
	})
}