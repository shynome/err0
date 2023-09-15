package try

type Error struct {
	error
}

func (e *Error) Unpack() error {
	return e.error
}

func To(e error) {
	if e == nil {
		return
	}
	if e, ok := e.(*Error); ok {
		panic(e)
	}
	panic(&Error{e})
}
