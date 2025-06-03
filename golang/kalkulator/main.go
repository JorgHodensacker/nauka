package main

import (
	"fmt"
	"os"

	"github.com/JorgHodensacker/nauka/golang/week1/pkg"
)

func main() {
	a, b, err := pkg.Input()
	if err != nil {
		fmt.Println(err)
	} else {
		result, err := pkg.Calc(os.Args[1], a, b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
