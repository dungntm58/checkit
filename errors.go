package checkit

type internalError struct {
	s string
}

func (e *internalError) Error() string {
	return e.s
}

func newInternalError(s string) error {
	return &internalError{
		s: s,
	}
}
