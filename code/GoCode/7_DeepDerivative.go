package main

import (
    "fmt"

    "k8s.io/apimachinery/pkg/api/equality"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    fmt.Println("******** struct instance comparison ********")

    person1 := Person{Name: "John", Age: 40}
    person2 := Person{Name: "John", Age: 40}
    person3 := Person{Name: "Jane", Age: 35}

    fmt.Println(equality.Semantic.DeepDerivative(person1, person2))
    fmt.Println(equality.Semantic.DeepDerivative(person1, person3))
}
