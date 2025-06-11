package pkg

import "errors"

func Calc(a float64, b float64, sign string) (float64, error) {
	switch sign {
	case "+":
		result, _ := Add(a, b)
		return result, nil
	case "-":
		result, err := Subtract(a, b)
		return result, err
	case "*":
		result, err := Multiply(a, b)
		return result, err
	case "/":
		result, err := Divide(a, b)
		return result, err
	default:
		return 0, errors.New("nieznana operacja")
	}
}
