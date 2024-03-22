package main

/*
Exercise 001:

Write a program which will receive a name and split into a map with the following config:

result["first_name"] = "Venerando"
result["second_name"] = "Lori"
result["third_name"] = "Endrizzi"

*/

import (
	"fmt"
	"reflect"
)

func Ex001(name string) map[string]string {
	result := make(map[string]string)

	//Write your solution here

	return result
}

func main() {
	fmt.Println("Exercise 01")

	want := make(map[string]string)
	want["first_name"] = "Venerando"
	want["second_name"] = "Lori"
	want["third_name"] = "Endrizzi"

	got := Ex001("Venerando Lori Endrizzi")

	if reflect.DeepEqual(got, want) {
		fmt.Println("is the same, Congrats!!")
	} else {
		fmt.Println("Values are different,  try again!")
		fmt.Println(got)
		fmt.Println(want)

	}
}
