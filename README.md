[![Build Status](https://travis-ci.org/mchirico/reflect.svg?branch=develop)](https://travis-ci.org/mchirico/reflect)

# reflect
Go Reflect Examples

[Good Examples](https://github.com/a8m/reflect-examples)



[Playground](https://play.golang.org/p/Gw0EzVKUZZn)

```go
package main

import (
	"fmt"
	"reflect"
)

type A struct{}

func (A) Fun0() { fmt.Println("fun0...called") }
func (A) Fun1() { fmt.Println("fun1...called") }

func Caller(i interface{}, f []string) {

	v := reflect.ValueOf(i)
	for _, val := range f {

		m := v.MethodByName(val)
		if m.Kind() != reflect.Func {
			fmt.Printf("Not Found: %s\n", val)
		} else {
			m.Call(nil)
		}
	}

}

func main() {

	s := []string{}
	s = append(s, "Fun0", "No Fun", "Fun1")
	Caller(A{}, s)
}

```


