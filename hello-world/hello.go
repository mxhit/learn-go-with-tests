package main

import "fmt"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const exclamationSuffix = "!"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name + exclamationSuffix
}

func main() {
	fmt.Println(Hello("Modo", "Spanish"))
}
