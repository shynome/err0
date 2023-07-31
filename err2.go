package err0

func NilThen(err *error, fn func()) {
	if *err == nil {
		fn()
	}
}

func To(ep *error) func(err error) error {
	return func(err error) error {
		*ep = err
		return err
	}
}
