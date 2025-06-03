package pkg

import (
	"errors"
	"os"
	"strconv"
)

func Input() (int, int, error) {
	if len(os.Args) < 3 {
		return 0, 0, errors.New("niepoprawne dane wejściowe")
	}
	a, err_a := strconv.Atoi(os.Args[2])
	b, err_b := strconv.Atoi(os.Args[3])
	if err_a != nil || err_b != nil {
		return 0, 0, errors.New("niepoprawne dane wejściowe")
	}
	return a, b, nil
}
