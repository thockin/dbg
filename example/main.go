// +build ignore

package main

import "github.com/thockin/dbg"

func main() {
	dbg.D.Printf("foo.1: %s", "bar")
	fn1(3)
	dbg.D.Printf("foo.2: %s", "qux")
}

func fn1(n int) {
	dbg.D.Printf("fn1(%d)", n)
	dbg.D.Push()
	defer dbg.D.Pop()
	dbg.D.Printf("scoped message")
	if n == 0 {
		return
	}
	fn1(n - 1)
}
