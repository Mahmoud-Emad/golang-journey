package Sum_all
import (
	"testing" 
	"reflect"
)

func TestSumAll(t *testing.T){
	got := SumAll([]int{1,3}, []int{0,9})
	want := []int{4,9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, Want %v", got, want)
	}
}