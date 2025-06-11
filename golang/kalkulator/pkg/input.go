package pkg

import "os"

func Input(a int, o int, b int) (string, string, string) {
	x := os.Args[a]
	y := os.Args[o]
	z := os.Args[b]

	return x, y, z
}
