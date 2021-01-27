package main

import "fmt"

func main() {
	fmt.Println(Hello("world", "English"))
}

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"
const spanishHolaPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const frenchBonjourPrefix = "Bonjour, "
const portugueseBonjourPrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchBonjourPrefix
	case spanish:
		prefix = spanishHolaPrefix
	case portuguese:
		prefix = portugueseBonjourPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
