package main

import (
	"fmt"
	"strings"
)

func main() {
	translations := make(map[string]string)

	addTranslation := func(word, translation string) error {
		if word == "" || translation == "" {
			return fmt.Errorf("Слово и его перевод не могут быть пустыми")
		}
		translations[word] = translation
		return nil
	}

	findTranslation := func(word string) string {
		if translation, ok := translations[word]; ok {
			return translation
		}
		return "Перевод не найден"
	}

	var inputWord, inputTranslation, choice string

	for {
		fmt.Print("Добавьте новый перевод (слово на одном языке и его перевод на другом язык, через пробел): ")
		_, err := fmt.Scanln(&inputWord, &inputTranslation)
		if err != nil {
			fmt.Println("Ошибка ввода")
			break
		}

		err = addTranslation(inputWord, inputTranslation)
		if err != nil {
			fmt.Println("Ошибка добавления перевода:", err)
			break
		}

		fmt.Print("Желаете добавить еще перевод? (y/n): ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) != "y" {
			break
		}
	}

	fmt.Println("\nПоиск перевода:")
	for {
		fmt.Print("Введите слово для поиска перевода: ")
		fmt.Scanln(&inputWord)
		translation := findTranslation(inputWord)
		fmt.Printf("Перевод слова '%s': %s\n", inputWord, translation)

		fmt.Print("Продолжить поиск? (y/n): ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) != "y" {
			break
		}
	}
}
