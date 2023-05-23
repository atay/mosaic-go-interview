package errors

type InvalidOperationError struct{}

func (e InvalidOperationError) Error() string {
	return "Invalid operation"
}

func (e InvalidOperationError) UserMessage() string {
	return e.Error()
}
