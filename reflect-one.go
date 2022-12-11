package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int64
}

func main() {
	foo := func() {
		fmt.Println("Hello")
	}
	foo()
	refObjTyp := reflect.TypeOf(foo)
	refObjVal := reflect.ValueOf(foo)
	refObjKnd := refObjTyp.Kind()
	refValKnd := refObjVal.Kind()
	if reflect.DeepEqual(refObjKnd, refValKnd) {
		fmt.Println("Same")
	} else {
		fmt.Println("Different")
	}
	fmt.Printf("Kind: %v\n", refObjTyp)
	fmt.Printf("Val: %v\n", refObjVal)
	fmt.Printf("Kind: %v\n", refObjKnd)
	fmt.Printf("Obj Kind: %v\n", refValKnd)
	refObjVal.Call(nil)

	// swapper
	i := 10
	j := 70
	sl := []int{i, 20, 30, 40, 50, 60, j}
	fu := reflect.Swapper(sl)
	fu(0, 6)
	fmt.Println(sl)
}
