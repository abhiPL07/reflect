package main

import (
	"fmt"
	"reflect"
)

type Marks struct {
	Physics float64
	Maths   float64
}

type Person struct {
	Name   string
	Age    int
	Origin struct {
		Country string
		City    string
	}
}

func main() {
	stud1 := Marks{
		Physics: 63,
		Maths:   91,
	}
	p1 := Person{
		Name: "Pal",
		Age:  30,
		Origin: struct {
			Country string
			City    string
		}{
			Country: "India",
			City:    "Kolkata",
		},
	}
	reflectItems(stud1)
	reflectItems(p1)
}

func reflectItems(it interface{}) error {
	typ := reflect.TypeOf(it)
	if typ.Kind() != reflect.Struct {
		fmt.Println("Not the expected value!")
	}
	fields := typ.NumField()
	fmt.Printf("Type : %v\nValue : %v\n", typ, fields)
	for i := 0; i < fields; i++ {
		tt := typ.Field(i)
		fmt.Printf("#: %v,\tName: %v,\tType: %v\n", i, tt.Name, tt.Type)
	}
	fmt.Println("________Next_________")
	return nil
}
