package hello_world

import "fmt"

const (
	english = "Enlgish"
	persian = "Persian"
	spanish = "Spanish"
	
	englishGreeting = "Hello, "
	persianGreeting = "Salam, "
	spanishGreeting = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishGreeting
	switch language {
	case persian:
		prefix = persianGreeting
	case spanish:
		prefix = spanishGreeting
	default:
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Soroush", english))
}
