package main

import "fmt"


func Hello(n string) string{
	return "Hello, super uper world and " + n
}

func main() {
	fmt.Println(Hello("Artus"))
}