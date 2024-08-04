package hello_world

import "fmt"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := "Hello, "
	switch language {
	case "Persian":
		prefix = "Salam, "
	case "Spanish":
		prefix = "Hola, "
	default:
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Soroush", "English"))
}
