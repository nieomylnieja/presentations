package main

import (
	"fmt"
)

type validationInterface[S any] interface { // CO
	Validate(s S) error // CO
} // CO
// CO
type Rule[T any] struct { // CO
	validate func(v T) error // CO
} // CO
// CO
type FieldGetter[T, S any] func(s S) T // CO
// CO
type FieldRules[T, S any] struct { // CO
	name   string                   // CO
	getter FieldGetter[T, S]        // CO
	rules  []validationInterface[T] // CO
} // CO
// CO
type StructValidator[S any] struct { // CO
	name   string                   // CO
	fields []validationInterface[S] // CO
} // CO
// CO
func (r Rule[T]) Validate(v T) error {
	return r.validate(v)
}

func (f FieldRules[T, S]) Validate(st S) error {
	value := f.getter(st)
	for _, rule := range f.rules {
		if err := rule.Validate(value); err != nil {
			return fmt.Errorf("invalid '%s' field value: %v", f.name, err)
		}
	}
	return nil
}

func (s StructValidator[S]) Validate(st S) error {
	for _, field := range s.fields {
		err := field.Validate(st)
		if err == nil {
			continue
		}
		return fmt.Errorf("validation for '%s' failed: %v", s.name, err)
	}
	return nil
}
