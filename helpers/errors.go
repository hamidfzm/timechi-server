package helpers

type ErrorAbort struct {
	message string
}

func (e *ErrorAbort) Error() string {
	return e.message
}
