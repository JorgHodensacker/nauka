package pkg

import (
	"fmt"
	"os"
)

// InitialCheck sprawdza, czy podano wystarczającą liczbę argumentów.
// x- liczba oczekiwanych argumentów
func InitialCheck(x int) {
	if len(os.Args) < x {
		fmt.Println("Nie podano wystarczającej liczby argumentów.")
		os.Exit(1)
	}
}
