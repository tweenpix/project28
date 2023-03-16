package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pkg/student"
)

func main() {
	studentMap := make(map[string]*student.Student)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите данные студента (имя возраст класс): ")
		if !scanner.Scan() {
			break
		}
		inputStr := scanner.Text()
		inputStr = strings.TrimSpace(inputStr)
		if len(inputStr) == 0 {
			continue
		}
		inputParts := strings.Split(inputStr, " ")
		if len(inputParts) != 3 {
			fmt.Println("Ошибка ввода данных")
			continue
		}
		name := inputParts[0]
		ageStr := inputParts[1]
		gradeStr := inputParts[2]
		student, err := student.New(name, ageStr, gradeStr)
		if err != nil {
			fmt.Println("Ошибка ввода данных")
			continue
		}
		studentMap[name] = student
	}
	fmt.Println("\nСтуденты из хранилища:")
	for name, student := range studentMap {
		fmt.Println(name, student.Age, student.Grade)
	}
}
