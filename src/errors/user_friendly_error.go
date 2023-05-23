package errors

type UserFriendlyError interface {
	error
	UserMessage() string
}
