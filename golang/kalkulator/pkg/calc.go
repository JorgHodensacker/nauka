package pkg

import "errors"

func Calc(ops string, a int, b int) (float64, error) {
	switch ops {
	case "dodaj":
		result, _ := Add(a, b)
		return result, nil
	case "odejmij":
		result, err := Subtract(a, b)
		return result, err
	case "pomnoz":
		result, err := Multiply(a, b)
		return result, err
	case "podziel":
		result, err := Divide(a, b)
		return result, err
	default:
		return 0, errors.New("nieznana operacja")
	}
}
