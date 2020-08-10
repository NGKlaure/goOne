package main

import (
	"fmt"

	"github.com/NGKlaure/goOne/calculation"
)

func main() {
	/* 	//fmt.Println("Hello world")
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

	   	fmt.Printf(greeting + ",Go!\n") */

	//variable declaration

	/* var x int
	x = 3
	var i, j = 100, "hello"
	fmt.Println("i and j is ", i, j)
	fmt.Println("x", x)
	*/

	//for loop
	/* var i int
	for i = 0; i < 5; i++ {
		fmt.Println(i)
	} */
	//if else statement
	/* var x int = 50
	if x < 10 {
		fmt.Println("x is less that 10")
	} else {
		fmt.Println("x is greater than 10")
	} */

	//switch case

	/* a, b := 1, 4

	switch a + b {
	case 1:
		fmt.Println("sum is 1")
	case 3:
		fmt.Println("sum is 3")
	default:
		fmt.Println("default")
	}
	*/

	//array systaxe
	/* var number [3]string //array of size 3
	direction := [...]int{1, 2, 3, 4, 5}
	number[0] = "one"
	number[1] = "twoo"
	number[2] = "tree"
	fmt.Println("the array is: ", number)
	fmt.Println("the array name direction is ", direction)
	fmt.Println("first elt is:", direction[0])
	fmt.Println("the lenght of direction is:", len(direction), "the capacity is:", cap(direction))
	*/

	//slice
	/*
		a := [5]string{"one", "two", "tree", "four", "five"}
		var b []string = a[1:4]
		fmt.Println("a: ", a)
		fmt.Println("b is: ", b) */

	//function call((calc))

	x, y := 15, 10
	sum, diff := calculation.Do_add(x, y) //used of package.Do_add is a method
	//in the package calculation and need to be with capital letter to be access anywere
	fmt.Println("the sum is: ", sum, "the difference is:", diff)

}

//function

/* func calc(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}
*/
