package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello world")
	var local string = os.Args[1]
	var greeting string = "hello"
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
