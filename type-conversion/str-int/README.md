# Learn Go With Tests

## How to Convert string to integer type in Go?

<p>

Like most modern languages, Golang includes strings as a built-in type. Let's take an example, you may have a string that contains a numeric value "100". However, because this value is represented as a string, you can't perform any mathematical calculations on it. You need to explicitly convert this string type into an integer type before you can perform any mathematical calculations on it. In order to convert string to integer type in Golang, you can use the following methods.

</p>


### `Atoi()`, `ParseInt()`, `fmt.Sscan`

* You can use the strconv package's Atoi() function to convert the string into an integer value. Atoi stands for ASCII to integer. The Atoi() function returns two values: the result of the conversion, and the error (if any).

* ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i. This function accepts a string parameter, convert it into a corresponding int type based on a base parameter. By default, it returns Int64 value.

* The fmt package provides sscan() function which scans string argument and store into variables. This function read the string with spaces and assign into consecutive Integer variables.


## Syntax
```go
func Atoi(s string) (int, error)
```

## Example
```go
package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}
```

## Output
```go
100 <nil> int
```

### For more examples see str-to-int.go file