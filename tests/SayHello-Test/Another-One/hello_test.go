package main
import "testing"

func HelloWithLangTest(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello in FR", func(t *testing.T) {
		got := helloWithLangs("Chris", "FR")
		want := "Oui, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in SP", func(t *testing.T) {
		got := helloWithLangs("Mahmoud", "SP")
		want := "Hola, Mahmoud"
		assertCorrectMessage(t, got, want)
	}) 
}