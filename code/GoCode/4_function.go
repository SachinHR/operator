package main

import (
    "fmt"
)

func main() {
    res, err := divide(10, 5)
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}