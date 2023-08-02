package err0

func NilThen(fn func()) {
	if r := recover(); r != nil {
		panic(r)
	}
	fn()
}
