package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int64
}

func main() {
	var p Person = Person{
		Name: "Lal",
		Age:  30,
	}
	res, err := JsonEncode(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}

func JsonEncode(v interface{}) ([]byte, error) {
	refObjVal := reflect.ValueOf(v)
	refObjTyp := reflect.TypeOf(v)
	buf := bytes.Buffer{}
	if refObjVal.Kind() != reflect.Struct {
		return buf.Bytes(), fmt.Errorf("unsupported type: %v", refObjVal.Kind())
	}
	pairs := []string{}
	for i := 0; i < refObjVal.NumField(); i++ {
		fieldVal := refObjVal.Field(i)
		fieldTyp := refObjTyp.Field(i)
		switch fieldVal.Kind() {
		case reflect.String:
			strVal := fieldVal.Interface().(string)
			pairs = append(pairs, `"`+string(fieldTyp.Name)+`": "`+strVal+`"`)
		case reflect.Int64:
			intVal := fieldVal.Interface().(int64)
			// pairs = append(pairs, `"`+string(fieldTyp.Name)+`": `+strconv.FormatInt(intVal, 10)+`"`)
			pairs = append(pairs, fmt.Sprintf("%s: %s", string(fieldTyp.Name), strconv.FormatInt(intVal, 10)))
		default:
			return buf.Bytes(), fmt.Errorf("unsupported type: %v", fieldVal.Kind())
		}
	}
	buf.WriteString("{")
	buf.WriteString(strings.Join(pairs, " "))
	buf.WriteString("}")
	return buf.Bytes(), nil
}
