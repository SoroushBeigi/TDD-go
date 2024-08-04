package hello_world

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	fmt.Println(Hello("Soroush"))
}
