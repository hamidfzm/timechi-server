package errors

type Authentication struct {
	message string
}

func (e *Authentication) Error() string {
	return e.message
}
