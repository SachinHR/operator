package main

import (
    "fmt"
    "reflect"
)

func main() {
    a := []int{1, 2, 3}
    b := []int{1, 2, 3}
    c := []int{1, 2, 4}

    fmt.Println(reflect.DeepEqual(a, b)) // true
    fmt.Println(reflect.DeepEqual(a, c)) // false
}