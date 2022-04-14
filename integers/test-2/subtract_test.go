package integers
import "testing"

func TestSubtract(t *testing.T){
	sub := Subtract(5, 2)
	expected := 3
	if sub != expected{
		t.Errorf("Expected %d but got %d.", expected, sub)
	}
}