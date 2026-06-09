package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Teacher struct {
	Name string `validate:"required"`
	Age  int    `validate:"required"`
}

func validate(v any) error {
	rv := reflect.ValueOf(v)                     // CO
	if rv.Kind() == reflect.Ptr && !rv.IsNil() { // CO
		rv = rv.Elem() // CO
	} // CO
	if rv.Kind() != reflect.Struct { // CO
		panic("only structs can be validated") // CO
	} // CO
	// ...
	rt := rv.Type()
	for i := range rv.NumField() {
		fieldValue := rv.Field(i)
		fieldType := rt.Field(i)
		if fieldType.Tag.Get("validate") == "required" && fieldValue.IsZero() {
			return errors.New(fieldType.Name + " field is required!")
		}
	}
	return nil
}

func main() {
	teacher := Teacher{Name: "Frank"}
	err := validate(teacher)
	fmt.Println("Error:", err)
}
