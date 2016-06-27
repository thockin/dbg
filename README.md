# dbg - Debugging prints

[![GoReportCard Widget]][GoReportCard]
[![GoDoc Widget]][GoDoc]

[GoDoc]: https://godoc.org/github.com/thockin/dbg
[GoDoc Widget]: https://godoc.org/github.com/thockin/dbg?status.svg
[GoReportCard]: https://goreportcard.com/report/github.com/thockin/dbg
[GoReportCard Widget]: https://goreportcard.com/badge/github.com/thockin/dbg

## Why?

Given the general lack of debugger support for Go (Delve is getting there), I
frequently find myself instrumenting code with random fmt.Printf() calls.  This
package is just a cleaner form of that.  There's nothing shocking or clever
here, just a chance to avoid some duplication of effort.

PRs welcome.
