package integers
import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T){
	sum := Add(2,2)
	expected := 4

	if sum != expected{
		t.Errorf("Expected %d but got %d", expected, sum)
	}	
}

// This will appear do the docmentation at http://localhost:6060/pkg/
// By adding this code the example will appear in the documentation inside godoc, making your code even more accessible.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}