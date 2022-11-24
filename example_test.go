package base62

import (
	"fmt"
)

func ExampleEncode() {
	fmt.Println(Encode(18446744073709551615))
	// Output: lYGhA16ahyf
}

func ExampleDecode() {
	fmt.Println(Decode("lYGhA16ahyf"))
	// Output: 18446744073709551615 <nil>
}
