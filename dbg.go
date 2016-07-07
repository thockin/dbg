// Package dbg offers some utility functions for debugging and debug-printing.
package dbg

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
)

// Printer is a debug handle, which hosts several methods for debug
// printing.
type Printer struct {
	logger       *log.Logger // output goes here
	indentString string      // string to print for each indent
	indent       int         // current indent level
}

// D is a global instance of a debug Printer.
var D = New(os.Stderr, "DBG: ")

// New allocates a new Printer.
func New(out io.Writer, prefix string) *Printer {
	return &Printer{
		logger:       log.New(out, prefix, 0),
		indentString: "_ ",
	}
}

// SetIndentString changes the string which is printed for each indent level.
// The default value is "_ ", which makes indents visually identifiable.
// Reasonable people may prefer plain whitespace or wider or narrower indents.
func (dbg *Printer) SetIndentString(str string) *Printer {
	dbg.indentString = str
	return dbg
}

// Push adds one indent level.
func (dbg *Printer) Push() {
	dbg.indent++
}

// Pop removes one indent level. If the indent level would become negative,
// this will simply exit.
func (dbg *Printer) Pop() {
	dbg.indent--
	if dbg.indent < 0 {
		fmt.Fprintf(os.Stderr, "Debug indent level became negative!\n\n")
		debug.PrintStack()
		os.Exit(42)
	}
}

// Printf is just like fmt.Printf, except that the user-provided value is
// prefixed with the Printer's prefix and zer or more indents.
func (dbg *Printer) Printf(f string, args ...interface{}) {
	s := ""
	for i := 0; i < dbg.indent; i++ {
		s += dbg.indentString
	}
	s += fmt.Sprintf(f, args...)
	dbg.logger.Println(s)
}
