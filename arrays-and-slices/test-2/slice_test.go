package slices
import "testing"

func TestSliceSum(t *testing.T){
	t.Run("Collection of any size", func(t *testing.T)  {
		numbers := []int{1,2,3} // Here you van pass any aitems you want
		got := Sum(numbers)
		expected := 6
		if got != expected{
			t.Errorf("Got %d, Expected %d, Given %d", got, expected, numbers)
		} 
	})
	
	// t.Run("Collection of fixed size", func(t *testing.T)  {
	// 	numbers := []int{1,2,3,4,5} // Here you van pass any aitems you want
	// 	got := Sum(numbers)
	// 	expected := 15
	// 	if got != expected{
	// 		t.Errorf("Got %d, Expected %d, Given %d", got, expected, numbers)
	// 	} 
	// })


}