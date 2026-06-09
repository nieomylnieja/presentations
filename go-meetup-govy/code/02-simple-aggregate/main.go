package main

import (
	"fmt"
	"slices"
	"strings"
)

type Teacher struct { // CO
	Name string // CO
	Age  int    // CO
} // CO
// CO

type multiError []string

func (m multiError) Error() string {
	return strings.Join(m, "; ")
}

func validateTeacher(teacher Teacher) error {
	var mErr multiError
	if len(teacher.Name) > 20 {
		mErr = append(mErr, "Teacher's name is too long!")
	}
	if !slices.Contains([]string{"John", "Eve"}, teacher.Name) {
		mErr = append(mErr, "we can only allow John and Eve to become teachers")
	}
	if teacher.Age < 20 {
		mErr = append(mErr, "Teacher is too young...")
	}
	return mErr
}

func main() {
	teacher := Teacher{Name: "Frank", Age: 18}
	err := validateTeacher(teacher)
	fmt.Println("Error:", err)
}
