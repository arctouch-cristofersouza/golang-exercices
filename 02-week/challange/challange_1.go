package main

import (
	"fmt"
	"reflect"
)

/*****

Fibonacci numbers form a sequence where each number is the sum of the two preceding ones,
usually starting with 0 and 1. The sequence goes: 0, 1, 1, 2, 3, 5, 8, 13, 21, and so forth.
wiki: https://en.wikipedia.org/wiki/Fibonacci_sequence

Write a func which a func( number int) which will return the sequence of fibonacci numbers in a string slice
******/

func fib(n int) []string {

	//return the correct value to test here, you can use another func to write you algorithm
	return []string{}
}

func main() {
	fmt.Println("Challenge")

	result_1 := fib(5)

	want := []string{"0", "1", "1", "2", "3", "5", "8", "13", "21", "34", "55", "89", "144", "233", "377", "610", "987", "1597", "2584", "4181"}

	if reflect.DeepEqual(result_1, want[:5]) {
		fmt.Println("Test passed! Fib with 5 numbers")
	} else {
		fmt.Println("Test failed! Fib with 5 numbers")
		fmt.Println("Want:")
		fmt.Println(want[:5])
		fmt.Println("Got:")
		fmt.Println(result_1)
		return
	}

	result_2 := fib(20)

	if reflect.DeepEqual(result_2, want) {
		fmt.Println("Test passed! Fib with 20 numbers")
	} else {
		fmt.Println("Test Failed! Fib with 20 numbers")
		fmt.Println("Want:")
		fmt.Println(want)
		fmt.Println("Got:")
		fmt.Println(result_2)
	}
}
