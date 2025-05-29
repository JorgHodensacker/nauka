package main

import (
	"fmt"

	"github.com/JorgHodensacker/nauka/golang/week1/pkg"
)

func main() {
	var a, b int

	a = 5
	b = 12

	result := pkg.Add(a, b)
	fmt.Printf(result)
}
