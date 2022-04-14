package main
import ("fmt")

// This is the first time i write this function,

// func Hello() string {
// 	return "Hello, world!"
// }

// Then we need to update this function to take at least one argument

// func Hello (name string) string{
// 	return "Hello, " + name
// }

// Refactor : Adding const value,
const englishHelloPrefix = "Hello, "

func Hello (name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main(){
	fmt.Println(Hello(""))
}