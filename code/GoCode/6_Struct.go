package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    var p1 Person
    var p2 Person

    p1.Name = "Alex"
    p1.Age = 45

    p2.Name = “Raj”
    p2.Age = 30

    fmt.Println("Name: ", p1.Name)
    fmt.Println("Age: ", p1.Age)

    fmt.Println("Name: ", p2.Name)
    fmt.Println("Age: ", p2.Age)
}
