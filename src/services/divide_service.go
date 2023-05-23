package services

import "errors"

func Divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}
