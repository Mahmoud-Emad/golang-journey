package Maps
import "testing"


// func TestSearch(t *testing.T){
// 	dictionary := map[string]string{"test": "This is just test value"}
// 	got := Search(dictionary, "test")
// 	want := "This is just test value"
// 	if got != want{
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// Refactor
// We can define our custom type Dictionary
// func TestSearchR(t *testing.T){
// 	dictionary := Dictionary{"test": "This is just test value"}
// 	got := dictionary.Search("test")
// 	want := "This is just test value"
// 	assertString(t, got, want)
// }

// func assertString(t testing.TB, got, want string){
// 	if got != want{
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// What if the key is not in our custom dictionary? Yes it's warrid to say that the program can continue without gettin any error.
// So lets define a new test case more complicated then the previous one.
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got,_ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	// t.Run("unknown word", func(t *testing.T) {
	// 	_, got := dictionary.Search("unknown")
	// 	assertError(t, got, NotFoundError())
	// })
}

func assertString(t testing.TB, got, want string){
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
