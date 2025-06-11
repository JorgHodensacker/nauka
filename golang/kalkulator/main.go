package main

import (
	"fmt"

	"github.com/JorgHodensacker/nauka/golang/week1/pkg"
)

// Zdefiniowanie stałych dla pozycji argumentów oddzielonych w linii poleceń
const (
	a             int = 1 // pierwszy argument znajduje się na pozycji 1
	operator      int = 2 // operator znajduje się na pozycji 2
	b             int = 3 // drugi argument znajduje się na pozycji 3
	supportedArgs int = 3 // oczekujemy trzech argumentów: a, operator, b
)

func main() {
	// Wstępne sprawdzenie, czy podano wystarczającą liczbę argumentów
	pkg.InitialCheck(supportedArgs)

	// Pobranie argumentów z linii poleceń
	x, y, z := pkg.Input(a, operator, b)

	// Zdefiniownie poprawnych argumentów i operatora
	a, b, sign := pkg.SyntaxCheck(x, y, z)

	// Wywołanie funkcji Calc z odpowiednimi argumentami
	result, err := pkg.Calc(a, b, sign)
	if err != nil {
		fmt.Println(err)
	}

	// Wyświetlenie wyniku
	fmt.Println("Wynik:", result)
}
