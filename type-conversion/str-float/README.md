# Learn Go With Tests

## How to Convert string to float type in Go?

<p>
Like most modern languages, Golang includes strings as a built-in type. Let's take an example, you may have a string that contains a numeric value "3.1415926535". However, because this value is represented as a string, you can't perform any mathematical calculations on it. You need to explicitly convert this string type into an float type before you can perform any mathematical calculations on it.
</p>

### `ParseFloat()`

* ParseFloat converts the string s to a floating-point number with the precision specified by bitSize: 32 for float32, or 64 for float64. When bitSize=32, the result still has type float64, but it will be convertible to float32 without changing its value.

## Syntax

``` func ParseFloat(s string, bitSize int) (float64, error) ```

## Example

```go
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

 s2 := "A-3141X"
 f2, err := strconv.ParseFloat(s2, 32)
 if err != nil {
  fmt.Println(err)
 } else {
  fmt.Println(f2, err, reflect.TypeOf(f2))
 }
}
```

## Output

```go
// 3.1415926535 <nil> float64
// -3.141 <nil> float64
// strconv.ParseFloat: parsing "A-3141X": invalid syntax
```

### For more examples see [str-to-float.go file](https://github.com/Mahmoud-Emad/Learn-Go-With-Tests/blob/master/type-conversion/str-float/str-to-float.go)
