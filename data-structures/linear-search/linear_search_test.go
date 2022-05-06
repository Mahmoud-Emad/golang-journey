package main
import "testing"



func TestLinearSearch(t *testing.T) {
	// Test 1
	data := []int{95,78,46,58,45,86,99,251,320}
	key := 58
	expected := true
	actual := linearSearch(data, key)
	if expected != actual {
		t.Errorf("Test 1: Expected %v, but got %v", expected, actual)
	}

	// Test 2
	data = []int{95,78,46,58,45,86,99,251,320}
	key = 123
	expected = false
	actual = linearSearch(data, key)
	if expected != actual {
		t.Errorf("Test 2: Expected %v, but got %v", expected, actual)
	}
}