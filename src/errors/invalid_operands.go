package errors

type InvalidOperandsError struct{}

func (e InvalidOperandsError) Error() string {
	return "Invalid operands"
}

func (e InvalidOperandsError) UserMessage() string {
	return e.Error()
}
