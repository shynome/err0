package err0

func NilThen(err *error, fn func()) {
	if *err == nil {
		fn()
	}
}
