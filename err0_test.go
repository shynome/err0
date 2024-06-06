package err0_test

import (
	"fmt"
	"testing"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

func TestThen(t *testing.T) {
	se := fmt.Errorf("hi, error!")
	t.Run("has error", func(t *testing.T) {
		var err error
		defer err0.Then(&err, nil, func() {
			if err != se {
				t.Error(err)
			}
		})
		defer err0.Then(&err, nil, func() {
			if err != se {
				t.Error(err)
			}
		})
		err0.Throw(se)
	})
	t.Run("no error", func(t *testing.T) {
		var err error
		defer err0.Then(&err, func() {
			if err != nil {
				t.Error(err)
			}
		}, func() { panic("wwwwwwwww") })
		defer err0.Then(&err, func() {
			if err != nil {
				t.Error(err)
			}
		}, func() { panic("wwwwwwwww") })
		try.To(nil)
	})
	t.Run("panic", func(t *testing.T) {
		var err error
		defer func() {
			if r := recover(); r != se {
				t.Error(r)
			}
		}()
		defer err0.Then(&err, nil, nil)
		panic(se)
	})
	t.Run("miao", func(t *testing.T) {
		se1 := fmt.Errorf("hi, error2!")
		var err error
		defer err0.Then(&err, nil, func() {
			panic("wwwwwwwww")
		})
		defer err0.Then(&err, nil, func() {
			if err != se1 {
				t.Error(err)
			}
			err = nil
		})
		defer err0.Then(&err, nil, func() {
			if err != se {
				t.Error(err)
			}
			try.To(se1)
		})
		try.To(se)
	})
}
