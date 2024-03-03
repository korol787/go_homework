package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Age    int
	Grades []int
}

func addStudent(students []Student, name string, age int, grades []int) []Student {
	student := Student{Name: name, Age: age, Grades: grades}
	students = append(students, student)
	return students
}

func listStudents(students []Student) {
	fmt.Println("Список студентов:")
	for i, student := range students {
		fmt.Printf("Студент №%d:\n", i+1)
		fmt.Printf("Имя: %s\n", student.Name)
		fmt.Printf("Возраст: %d\n", student.Age)
		fmt.Printf("Оценки: %v\n", student.Grades)
		fmt.Println("---------------------------")
	}
}

func main() {
	var students []Student

	var name string
	var ageStr, gradeStr string
	var age int
	var grades []int

	for {
		fmt.Println("Добавление информации о студенте:")

		fmt.Print("Имя студента: ")
		fmt.Scanln(&name)

		fmt.Print("Возраст студента: ")
		_, err := fmt.Scanln(&ageStr)
		if err != nil {
			fmt.Println("Ошибка ввода возраста.")
			break
		}
		age, err = strconv.Atoi(ageStr)
		if err != nil {
			fmt.Println("Ошибка преобразования возраста в число.")
			break
		}
		if age < 0 || age > 100 {
			fmt.Println("Некорректный возраст.")
			break
		}

		fmt.Print("Оценки студента (через запятую): ")
		_, err = fmt.Scanln(&gradeStr)
		if err != nil {
			fmt.Println("Ошибка ввода оценок.")
			break
		}
		gradeStr = strings.ReplaceAll(gradeStr, " ", "")
		gradeArr := strings.Split(gradeStr, ",")
		for _, g := range gradeArr {
			var grade int
			grade, err = strconv.Atoi(g)
			if err != nil {
				fmt.Println("Ошибка преобразования оценки в число.")
				break
			}
			if grade < 1 || grade > 5 {
				fmt.Println("Некорректная оценка.")
				break
			}
			grades = append(grades, grade)
		}

		students = addStudent(students, name, age, grades)
		grades = nil

		var choice string
		fmt.Print("Желаете добавить еще студента? (y/n): ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) != "y" {
			break
		}
	}

	listStudents(students)
}
