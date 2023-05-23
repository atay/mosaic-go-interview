package services

type CannotDivideByZeroError struct{}

func (e CannotDivideByZeroError) Error() string {
	return "Cannot divide by zero"
}

func Divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, CannotDivideByZeroError{}
	}
	return x / y, nil
}
