package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p1 := &Person{Name: "John", Age: 30}
    p2 := &Person{Name: "Alex", Age: 29}

    result := compareAge(p1, p2)
    fmt.Println(result)
}

func compareAge(p1, p2 *Person) string {
    if p1.Age == p2.Age {
        return "equal"
    }
    return "Notequal"
}
