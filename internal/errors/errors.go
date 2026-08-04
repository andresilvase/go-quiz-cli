package errors

type ExpectInteger struct {
	Message string
}

func (e ExpectInteger) Error() string {
	return e.Message
}
