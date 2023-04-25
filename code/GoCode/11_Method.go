package main

import "fmt"

type rectangle struct {
    width  int
    height int
}

func (r rectangle) area() int {
    return r.width * r.height
}

func main() {
    r := rectangle{width: 10, height: 5}

    fmt.Println("The area of the rectangle is", r.area())
}