package Maps
import "testing"


func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	got, err := dictionary.Search("test")
	expected := "this is just a test"

	if got != expected {
		t.Errorf("There are an error in adding method.")
	}

	if err != nil {
		t.Fatal("Should find added word.", err)
	}
}

// Refactor
// There isn't much to refactor in our implementation but the test could use a little simplification.

func TestAddR(t *testing.T) {
	dictionary := Dictionary{}
	key := "test"
	value := "this is a test value"
	dictionary.Add(key, value)
	assertDefinition(t, dictionary, key, value)
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if got != value {
		t.Errorf("got %q want %q", got, value)
	}

	if err != nil {
		t.Fatal("Should find added word.", err)
	}
}