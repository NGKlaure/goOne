package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello world")
	var local string
	var greeting string = "hello"

	fmt.Println("enter the code language")
	fmt.Scanln(&local)

	if local == "en" {
		greeting = "hello"
	} else if local == "fr" {
		greeting = "bonjour"
	} else if local == "es" {
		greeting = "hola"
	} else if local == "de" {
		greeting = "gutengtag"
	} else {
		greeting = "yo"
	}

	fmt.Printf(greeting + ",Go!\n")

}
