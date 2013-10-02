# Gospel
BDD-style testing library for Golang.

## Install
```
$ go get github.com/r7kamura/gospel
```

## Usage
```go
package main

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe(t *testing.T) {
	Describe(t, "gospel.Expectation#ToEqual", func(context Context, it It) {
		context("with 1 & 1", func() {
			it("compares integers with ==", func(expect Expect) {
				expect(1).ToEqual(1)
			})
		})
		context("with `1` & `1`", func() {
			it("compares strings with ==", func(expect Expect) {
				expect("1").ToEqual("1")
			})
		})
	})

	Describe(t, "gospel.Expectation#ToNotEqual", func(context Context, it It) {
		context("with 1 & 2", func() {
			it("compares integers with !=", func(expect Expect) {
				expect(1).ToNotEqual(2)
			})
		})
		context("with `1` & `2`", func() {
			it("compares strings with !=", func(expect Expect) {
				expect("1").ToNotEqual("2")
			})
		})
	})
}
```

```
$ go test
gospel.Expectation#ToEqual
        with 1 & 1
                compares integers with ==
        with `1` & `1`
                compares strings with ==
gospel.Expectation#ToNotEqual
        with 1 & 2
                compares integers with !=
        with `1` & `2`
                compares strings with !=
PASS
ok      _/Users/r7kamura/gospel     0.023s
```
