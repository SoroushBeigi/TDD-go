package iteration

import "fmt"

func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

// ExampleRepeat demonstrates the use of the Repeat function.
func ExampleRepeat() {
	result1 := Repeat("a", 3)
	fmt.Println(result1)
	//Output: aaa

	result2 := Repeat("a", 10)
	fmt.Println(result2)
	//Output: aaaaaaaaaa
}
