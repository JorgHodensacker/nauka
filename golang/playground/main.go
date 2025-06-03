package main

func main() {
	type error interface {
		Error() string
	}
}
