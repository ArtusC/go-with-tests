package hello

import "fmt"

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

const english = "English"
const englishHelloPrefix = "Hello, super world and "

const portugues = "Portugues"
const portuguesHelloPrefix = "Olá, super mundo e "

func Hello(name, language string) string {
	if name == "" {
		return englishHelloPrefix + "anything!"
	}

	return greetingPrefix(language) + name + "!"

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case portugues:
		prefix = portuguesHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Artus", "Portugues"))
	fmt.Println(Hello("Artus", "English"))
	fmt.Println(Hello("", "Portugues"))
	fmt.Println(Hello("", "English"))
	fmt.Println(Hello("", ""))
}
