package arrays
import (
	"testing"
)

func TestSum(t *testing.T){
	numbers := [5] int{1,2,3,4,5}
	got := Sum(numbers)
	expected := 15
	if got !=  expected{
		t.Errorf("Got %d expected %d given. %v", got, expected, numbers)
	}
}