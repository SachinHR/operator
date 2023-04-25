package main

import "fmt"

func main() {
	if x := 10; x < 0 {
		fmt.Println("x is negative")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is positive")
	 }
}