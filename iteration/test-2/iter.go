package iteration
import "fmt"

func Repeat(character string, num int) string {
	// caller can specify how many times the character is repeated
	var repeated string
	for i := 0; i < num; i++ {
		fmt.Println(character, i)
		repeated += character
	}
	return repeated
}