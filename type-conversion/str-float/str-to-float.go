// Look at the readmefile for more information.

package main

import (
 "fmt"
 "reflect"
 "strconv"
)

func main() {
	s := "3.1415926535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))

	s1 := "-3.141"
	f1, err := strconv.ParseFloat(s1, 8)
	fmt.Println(f1, err, reflect.TypeOf(f1))

	//  This will return an error because the string is not a float number.
	s2 := "A-3141X"
	f2, err := strconv.ParseFloat(s2, 32)

	if err != nil {
		// To check if an error occurred during the conversion, check if the err variable contains a nil value.
		fmt.Println(err)
	} else {
		fmt.Println(f2, err, reflect.TypeOf(f2))
	}
}