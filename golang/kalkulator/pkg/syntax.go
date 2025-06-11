package pkg

import (
	"fmt"
	"os"
	"strconv"
)

// SyntaxCheck sprawdza poprawność składniową argumentów wejściowych, oczekując:
// x - pierwszy argument (liczba)
// y - operator (jeden z: "+", "-", "*", "/")
// z - drugi argument (liczba)
func SyntaxCheck(x string, y string, z string) (float64, float64, string) {
	// Sprawdzenie, czy pierwszy argument jest liczbą
	a, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Println("Pierwszy argument nie jest liczbą:", 1)
		os.Exit(1)
	}

	// Sprawdzenie, czy drugi argument jest liczbą
	b, err := strconv.ParseFloat(z, 64)
	if err != nil {
		fmt.Println("Drugi argument nie jest liczbą:", 3)
		os.Exit(1)
	}

	// Sprawdzenie, czy operator jest jednym z dozwolonych
	sign := y
	switch sign {
	case "+", "-", "*", "/":
	default:
		fmt.Println("Nieznany operator:", sign)
		os.Exit(1)
	}

	// Zwrócenie poprawnych argumentów i operatora
	return a, b, sign
}
