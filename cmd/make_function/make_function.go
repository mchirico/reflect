
// Ref: https://github.com/golang/go/issues/20189

package main

import (
	"fmt"
)

type S struct {
	s string
}

func (s *S) makeFunction(name string) func() {
	return func() {
		fmt.Printf("%s %s", s.s, name)
	}
}
func main() {
	s := S{s: "svalue:"}
	f := s.makeFunction("Ilja")
	f()
}

