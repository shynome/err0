package err0

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lainio/err2"
	"github.com/lainio/err2/assert"
	"github.com/lainio/err2/try"
)

func TestBind(t *testing.T) {

	t.Run("0", func(t *testing.T) {
		var terr = fmt.Errorf("some err")
		defer err2.Catch(func(err error) {
			if errors.Is(err, terr) {
				return
			}
			t.Error(err)
		})

		var a = 0
		defer func() {
			assert.Equal(a, 0)
		}()

		var err error
		defer NilThen(&err, func() { a = 1 })
		try.Out(terr).Handle(To(&err))
	})

	t.Run("1", func(t *testing.T) {
		defer err2.Catch(func(err error) { t.Error(err) })

		var a = 0
		defer func() {
			assert.Equal(a, 1)
		}()

		var err error
		defer NilThen(&err, func() { a = 1 })
		try.Out(nil).Handle(To(&err))
	})
}
