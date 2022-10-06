package main

import "fmt"

// Code to run "hello world" using constant and function

const helloPrefix = "Hello, super world and "

func Hello(name string) string{
	if name != "" {
		return helloPrefix + name + "!"
	} else {
		return helloPrefix + "anything!"
	}
}

func main() {
	fmt.Println(Hello("Artus"))
}


