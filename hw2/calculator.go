package main

import "fmt"

var dohod, food, transport, living, net_income float32

func main() {
	fmt.Print("Введите ваш месячный доход: ")
	fmt.Scanln(&dohod)
	fmt.Print("Введите расходы на еду: ")
	fmt.Scanln(&food)
	fmt.Print("Введите расходы на транспорт: ")
	fmt.Scanln(&transport)
	fmt.Print("Введите расходы на жилье: ")
	fmt.Scanln(&living)

	net_income := dohod - food - transport - living

	fmt.Printf("Чистый доход: %v \n", net_income)
	if food/dohod >= 0.3 {
		fmt.Printf("Внимание: расходы на еду составляют %.0f%% от вашего дохода. Рекомендуется снизить траты в этой категории.", food/dohod*100)
	} else if transport/dohod >= 0.3 {
		fmt.Printf("Внимание: расходы на транспорт составляют %.0f%% от вашего дохода. Рекомендуется снизить траты в этой категории.", transport/dohod*100)
	} else if living/dohod >= 0.3 {
		fmt.Printf("Внимание: расходы на жильё составляют %.0f%% от вашего дохода. Рекомендуется снизить траты в этой категории.", living/dohod*100)
	} else {
		fmt.Print("Ваши расходы сбалансированы по статьям")
	}
}
