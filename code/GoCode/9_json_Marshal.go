package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:”PersonName"`
    Age  int    `json:”PersonAge"`
}

func main() {
    p1 := Person{Name: "John", Age: 30}
    jsonData, err := json.Marshal(p1)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("JSON data:", string(jsonData))
}
