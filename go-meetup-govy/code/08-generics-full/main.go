package main

import (
	"errors"
	"fmt"
)

type Teacher struct { // CO
	Name string // CO
	Age  int    // CO
} // CO
// CO
type validationInterface[S any] interface { // CO
	Validate(s S) error // CO
} // CO
// CO
func NewRule[T any](validate func(v T) error) Rule[T] { // CO
	return Rule[T]{validate: validate} // CO
} // CO
// CO
type Rule[T any] struct { // CO
	validate func(v T) error // CO
} // CO
// CO
func (r Rule[T]) Validate(v T) error { return r.validate(v) } // CO
// CO
type FieldGetter[T, S any] func(s S) T // CO
// CO
func NewFieldRules[T, S any]( // CO
	name string, // CO
	getter FieldGetter[T, S], // CO
	rules ...validationInterface[T], // CO
) FieldRules[T, S] { // CO
	return FieldRules[T, S]{ // CO
		name:   name,   // CO
		getter: getter, // CO
		rules:  rules,  // CO
	} // CO
} // CO
// CO
type FieldRules[T, S any] struct { // CO
	name   string                   // CO
	getter FieldGetter[T, S]        // CO
	rules  []validationInterface[T] // CO
} // CO
// CO
func (f FieldRules[T, S]) Validate(st S) error { // CO
	value := f.getter(st)          // CO
	for _, rule := range f.rules { // CO
		if err := rule.Validate(value); err != nil { // CO
			return fmt.Errorf("invalid '%s' field value: %v", f.name, err) // CO
		} // CO
	} // CO
	return nil // CO
} // CO
// CO
func NewStructValidator[S any](name string, fields ...validationInterface[S]) StructValidator[S] { // CO
	return StructValidator[S]{name: name, fields: fields} // CO
} // CO
// CO
type StructValidator[S any] struct { // CO
	name   string                   // CO
	fields []validationInterface[S] // CO
} // CO
// CO
func (s StructValidator[S]) Validate(st S) error { // CO
	for _, field := range s.fields { // CO
		err := field.Validate(st) // CO
		if err == nil {           // CO
			continue // CO
		} // CO
		return fmt.Errorf("validation for '%s' failed: %v", s.name, err) // CO
	} // CO
	return nil // CO
} // CO
// CO
var (
	nameRule = NewRule(func(v string) error {
		if len(v) > 20 {
			return errors.New("name is too long!")
		}
		return nil
	})
	ageRule = NewRule(func(v int) error {
		if v < 20 {
			return errors.New("too young...")
		}
		return nil
	})
)

var teacherValidator = NewStructValidator(
	"Teacher",
	NewFieldRules("Name", func(t Teacher) string { return t.Name }, nameRule),
	NewFieldRules("Age", func(t Teacher) int { return t.Age }, ageRule),
)

func main() {
	teacher := Teacher{Name: "Frank", Age: 18}
	err := teacherValidator.Validate(teacher)
	fmt.Println("Error:", err)
}
