# Base62

Encode uint64 into a base62 string value, or decode a base62 string into an int64 value.

## Usage

``` go
package main

import (
    "fmt"
    "github.com/divinerapier/base62"
)

func ExampleEncode() {
 fmt.Println(Encode(18446744073709551615))
 // Output: lYGhA16ahyf
}

func ExampleDecode() {
 fmt.Println(Decode("lYGhA16ahyf"))
 // Output: 18446744073709551615 <nil>
}
```
