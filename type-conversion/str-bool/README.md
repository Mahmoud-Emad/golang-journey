# Learn Go With Tests

## How to convert String to Boolean Data Type Conversion in Go?

<p>
Like most modern languages, Golang includes strings as a built-in type. Let's take an example, you may have a string that contains a boolean value "true". However, because this value is represented as a string, you can't perform any operation on it. You need to explicitly convert this string type into an boolean type before you can perform any operation on it.
</p>

### String to Boolean Conversion

* Package strconv is imported to perform conversions to and from string.ParseBool returns the boolean value represented by the string. It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.

## Example

```go
package main

import (
 "fmt"
 "strconv"
)

func main() {
 s1 := "true"
 b1, _ := strconv.ParseBool(s1)
 fmt.Printf("%T, %v\n", b1, b1)

 s2 := "t"
 b2, _ := strconv.ParseBool(s2)
 fmt.Printf("%T, %v\n", b2, b2)

 s3 := "0"
 b3, _ := strconv.ParseBool(s3)
 fmt.Printf("%T, %v\n", b3, b3)

 s4 := "F"
 b4, _ := strconv.ParseBool(s4)
 fmt.Printf("%T, %v\n", b4, b4)
}
```

## Output

```go
// bool, true
// bool, true
// bool, false
// bool, false
```
