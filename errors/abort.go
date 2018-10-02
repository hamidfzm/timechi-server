package errors

type Abort struct {
	message string
}

func (e *Abort) Error() string {
	return e.message
}
