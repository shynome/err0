## Introduction

The Repo API is from <https://github.com/lainio/err2/tree/v0.9.1#error-checks>

I like err2 old api, so I create this.

## API

```go
package main

import (
	"fmt"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

func main() {
	var err error
	defer err0.Then(&err, func() {
		fmt.Println("no error", err)
	}, func() {
		fmt.Println("has error", err)
	})
	try.To(fmt.Errorf("hi, error!"))
}

```
