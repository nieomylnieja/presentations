package main

import (
	_ "fmt"
)

type validationInterface[S any] interface {
	Validate(s S) error
}

func NewRule[T any](validate func(v T) error) Rule[T] {
	return Rule[T]{validate: validate}
}

type Rule[T any] struct {
	validate func(v T) error
}

type FieldGetter[T, S any] func(s S) T

func NewFieldRules[T, S any](
	name string,
	getter FieldGetter[T, S],
	rules ...validationInterface[T],
) FieldRules[T, S] {
	return FieldRules[T, S]{
		name:   name,
		getter: getter,
		rules:  rules,
	}
}

type FieldRules[T, S any] struct {
	name   string
	getter FieldGetter[T, S]
	rules  []validationInterface[T]
}

func NewStructValidator[S any](name string, fields ...validationInterface[S]) StructValidator[S] {
	return StructValidator[S]{name: name, fields: fields}
}

type StructValidator[S any] struct {
	name   string
	fields []validationInterface[S]
}
