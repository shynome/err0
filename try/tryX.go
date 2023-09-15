package try

func To1[T1 any](t1 T1, e error) T1 {
	To(e)
	return t1
}

func To2[T1 any, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	To(e)
	return t1, t2
}

func To3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, e error) (T1, T2, T3) {
	To(e)
	return t1, t2, t3
}
