package main

import (
	"fmt"
	"reflect"
)

type T func(int64, int64) int64

func main() {
	t := reflect.TypeOf(T(nil))

	mul := reflect.MakeFunc(t, func(args []reflect.Value) []reflect.Value {
		a := args[0].Int()
		b := args[1].Int()
		return []reflect.Value{reflect.ValueOf(a+b)}
	})
	fn, ok := mul.Interface().(T)
	if !ok {
		return
	}
	fmt.Println(fn(2,3))
}
