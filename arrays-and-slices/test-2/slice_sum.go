package slices
// import "fmt"

func Sum(numbers []int) int{
	result := 0
	for _, number := range numbers{
		result += number
	}
	// fmt.Printf("[%d %d]\n", i, number) // X here the index of n in array
	return result
}