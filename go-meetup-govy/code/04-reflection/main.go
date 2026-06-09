package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Teacher struct { // CO
	Name string // CO
	Age  int // CO
} // CO
// CO
type Rule func(v any) error

var teacherRules = map[string]Rule{
	"Name": func(v any) error {
		if len(v.(string)) > 20 {
			return errors.New("Teacher's name is too long!")
		}
		return nil
	},
	"Age": func(v any) error {
		if v.(int) < 20 {
			return errors.New("Teacher is too young...")
		}
		return nil
	},
}

func validate(v any, rules map[string]Rule) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		panic("only structs can be validated")
	}
	for fieldName, rule := range rules {
		field := rv.FieldByName(fieldName)
		if err := rule(field.Interface()); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	teacher := Teacher{Name: "Frank", Age: 18}
	err := validate(teacher, teacherRules)
	fmt.Println("Error:", err)
}
