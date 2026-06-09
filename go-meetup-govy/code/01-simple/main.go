package main

import (
	"errors"
	"fmt"
	"slices"
)

type Teacher struct {
	Name string
	Age  int
}

func validateTeacher(teacher Teacher) error {
	if len(teacher.Name) > 20 {
		return errors.New("Teacher's name is too long!")
	}
	if !slices.Contains([]string{"John", "Eve"}, teacher.Name) {
		return errors.New("we can only allow John and Eve to become teachers")
	}
	if teacher.Age < 20 {
		return errors.New("Teacher is too young...")
	}
	return nil
}

func main() {
	teacher := Teacher{Name: "Frank", Age: 18}
	err := validateTeacher(teacher)
	fmt.Println("Error:", err)
}
