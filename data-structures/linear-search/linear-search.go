package main
import (
	"fmt"
)


func linearSearch(data []int, key int) bool {
	// Check if the key is in the data
	for _, item := range data {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
    items := []int{95,78,46,58,45,86,99,251,320}
    fmt.Println(linearSearch(items,58))
}