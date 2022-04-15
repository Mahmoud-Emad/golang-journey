package arrays


// import "fmt"
// func Sum(numbers [5]int)int{
// 	var current int
// 	for i := 0; i < 5; i++{
// 		current += numbers[i]
// 	}
// 	fmt.Println(current)
// 	return current
// }

// Refactor
import "fmt"
func Sum(numbers [5]int) int {
	sum := 0
	for x, number := range numbers {
		fmt.Printf("[%d %d]\n", x, number) // X here the index of n in array
		sum += number 
	}
	return sum
}