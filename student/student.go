package student

import "strconv"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func New(name string, ageStr string, gradeStr string) (*Student, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, err
	}
	grade, err := strconv.Atoi(gradeStr)
	if err != nil {
		return nil, err
	}
	return &Student{name, age, grade}, nil
}
