package err0

import "github.com/shynome/err0/try"

func Throw(err error) {
	try.To(err)
}
