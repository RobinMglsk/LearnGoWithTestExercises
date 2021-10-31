package main

import "fmt"

const spanishLocale = "es"
const frenchLocale = "fr"
const dutchLocale = "nl"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const dutchHelloPrefix = "Hallo, "

func Hello(name string, locale string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(locale) + name
}

func greetingPrefix(locale string) (prefix string) {
	switch locale {
	case spanishLocale:
		prefix = spanishHelloPrefix
	case frenchLocale:
		prefix = frenchHelloPrefix
	case dutchLocale:
		prefix = dutchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", "en"))
}
