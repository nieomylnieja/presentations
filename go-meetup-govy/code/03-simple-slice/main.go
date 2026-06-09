package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type multiError []string // CO
// CO
func (m multiError) Error() string { // CO
	return strings.Join(m, "; ") // CO
} // CO
// CO
type Teacher struct {
	Name string // CO
	Age  int    // CO
	// ...
	Students []Student
}

type Student struct {
	Index string
}

func validateTeacher(teacher Teacher) error {
	var mErr multiError         // CO
	if len(teacher.Name) > 20 { // CO
		mErr = append(mErr, "Teacher's name is too long!") // CO
	} // CO
	if !slices.Contains([]string{"John", "Eve"}, teacher.Name) { // CO
		mErr = append(mErr, "we can only allow John and Eve to become teachers") // CO
	} // CO
	if teacher.Age < 20 { // CO
		mErr = append(mErr, "Teacher is too young...") // CO
	} // CO
	// ...
	for i, student := range teacher.Students {
		err := validateStudent(student)
		if err == nil {
			continue
		}
		mErr = append(mErr, fmt.Sprintf("student[%d]: %v", i, err.Error()))
	}
	return mErr
}

func validateStudent(student Student) error {
	intIdx, err := strconv.Atoi(student.Index)
	switch {
	case err != nil:
		return fmt.Errorf("invalid index: %w", err)
	case intIdx == 0:
		return errors.New("Index must be non-zero")
	default:
		return nil
	}
}

func main() {
	teacher := Teacher{Name: "John", Age: 22, Students: []Student{
		{Index: "102"},
		{Index: "00"},
	}}
	err := validateTeacher(teacher)
	fmt.Println("Error:", err)
}
