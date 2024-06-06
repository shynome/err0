package err0

import "github.com/shynome/err0/try"

// Then only handle try.ToX panic generate error(try.Error)
func Then(err *error, onok func(), onerr func()) {
	e := recover()
	if e != nil {
		if te, ok := e.(*try.Error); !ok {
			panic(e)
		} else {
			*err = te.Unpack()
		}
	}
	switch {
	case *err == nil && onok != nil:
		onok()
	case *err != nil && onerr != nil:
		onerr()
	}
}

var Throw = try.To
