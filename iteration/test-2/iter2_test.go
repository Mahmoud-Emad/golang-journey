package iteration
import (
	"fmt"
	"testing"
)



func TestRepeat(t *testing.T){
	num := 10
	repeated := Repeat("a", num)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}



func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat(){
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}