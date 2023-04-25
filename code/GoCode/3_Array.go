package main

import (
	"fmt"
)

func main() {
	// Initialize a string slice
	strList := []string{"apple", "banana", "cherry"}

	// Use the append function to add elements to the slice
	strList = append(strList, "date", "mongo")

	fmt.Println(strList)
}