package main

/*
Exercise 001:

Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5,
 between 2000 and 3200 (both included).
 The numbers obtained should be printed in a comma-separated sequence on a single line.
*/

import (
	"fmt"
)

func main() {
	fmt.Println("Exercise")

	res := Ex001(2000, 3200)

	want := "112,119,126,133,147,154,161,168,182,189,196"

	if res == want {
		fmt.Println("You got it, all numbers are correct")
	} else {
		fmt.Println("Values are different,  try again!")
		fmt.Println("Want:")
		fmt.Println(want)
		fmt.Println("Got:")
		fmt.Println(res)
	}

}

// Ex001 returns a slice of numbers
func Ex001(low, high int) string {

	//Write your solution here

	return ""
}
