package main

import "fmt"

// Code to run "hello world" using constant and function

const english = "English"
const englishHelloPrefix = "Hello, super world and "

const portugues = "Portugues"
const portuguesHelloPrefix = "Olá, super mundo e "

func Hello(name, language string) string{
	if language == english {
		if name != "" {
			return englishHelloPrefix + name + "!"
		} else {
			return englishHelloPrefix + "anything!"
		}
	}

	if language == portugues {
		if name != "" {
			return portuguesHelloPrefix + name + "!"
		} else {
			return portuguesHelloPrefix + "ninguém!"
		}
	}

	return ""
	
}

func main() {
	fmt.Println(Hello("Artus", "Portugues"))
}


