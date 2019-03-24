package main

import (
	"fmt"
	"reflect"
)

type T func(int64, int64) int64

type S struct {
	s string
}

func (s *S)SetS(t string) string {
	s.s = t
	return s.s
}


func E(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	fmt.Printf("T= %T\n",i)
	return t

}


func D(i interface{}) func(int64, int64) int64 {
	t := reflect.TypeOf(i)
	mul := reflect.MakeFunc(t, func(args []reflect.Value) []reflect.Value {
		a := args[0].Int()
		b := args[1].Int()
		return []reflect.Value{reflect.ValueOf(a + b)}
	})
	fn, ok := mul.Interface().(T)
	if !ok {
		return nil
	}
	return fn

}

func main() {

	s := S{}
	s.SetS("one")
	E(s)




	fn2 := D(T(nil))
	if fn2 != nil {
		fmt.Println(fn2(2, 4))
	}

}
