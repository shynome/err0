package err0

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lainio/err2"
	"github.com/lainio/err2/assert"
	"github.com/lainio/err2/try"
)

func TestErrBind(t *testing.T) {

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

		defer NilThen(func() { a = 1 })

		try.To(terr)
	})

	t.Run("1", func(t *testing.T) {
		defer err2.Catch(func(err error) { t.Error(err) })

		var a = 0
		defer func() {
			assert.Equal(a, 1)
		}()

		defer NilThen(func() { a = 1 })

		try.To(nil)
	})
}
