package pkg

import "errors"

func Add(a float64, b float64) (float64, error) {
	sum := float64(a + b)
	return sum, nil
}

func Subtract(a float64, b float64) (float64, error) {
	diff := float64(a - b)
	return diff, nil
}

func Multiply(a float64, b float64) (float64, error) {
	product := float64(a * b)
	return product, nil
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("dzielenie przez zero")
	}
	quotient := float64(a) / float64(b)
	return quotient, nil
}
