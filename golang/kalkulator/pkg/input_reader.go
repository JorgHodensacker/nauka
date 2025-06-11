package pkg

import "os"

func Input(p1 int, p2 int, p3 int) (string, string, string) {
	a := os.Args[p1]
	o := os.Args[p2]
	b := os.Args[p3]

	return a, o, b
}
