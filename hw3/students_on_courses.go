package main

import (
	"fmt"
)

func main() {
	courseStudents := make(map[string]int)

	addStudentToCourse := func(course string) error {
		if course == "" {
			return fmt.Errorf("Ошибка: название курса не может быть пустым")
		}
		courseStudents[course]++
		return nil
	}

	printTotalStudentsPerCourse := func() {
		fmt.Println("Общее количество студентов по каждому курсу:")
		for course, students := range courseStudents {
			fmt.Printf("Курс %s: %d студент(ов)\n", course, students)
		}
	}

	var course, choice string
	var err error

	for {
		fmt.Print("Введите название курса: ")
		_, err = fmt.Scanln(&course)
		if err != nil {
			fmt.Println("Ошибка ввода.")
			break
		}

		err = addStudentToCourse(course)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Print("Желаете добавить студента на другой курс? (y/n): ")
		_, err = fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода.")
			break
		}

		if choice != "y" {
			break
		}
	}

	printTotalStudentsPerCourse()
}
