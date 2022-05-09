// Look at the readmefile for more information.
package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	strVar := "100"
	intVar1, err := strconv.Atoi(strVar)
	fmt.Println(intVar1, err, reflect.TypeOf(intVar1))
	// Output: 100 <nil> int
	
	// Another example:
	intVar2, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	intVar2, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	intVar2, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	intVar2, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	// Output:
	// 100 <nil> int64
	// 100 <nil> int64
	// 100 <nil> int64
	// 100 <nil> int64

	// Another example:
	intValue3 := 0
	_, err = fmt.Sscan(strVar, &intValue3)
	fmt.Println(intValue3, err, reflect.TypeOf(intValue3))
	// Output:
	// 100 <nil> int
}