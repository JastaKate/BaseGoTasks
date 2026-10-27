package main

import "fmt"

func main() {
	fmt.Println("Задача 1. Расчёт стоимости билетов")
	price := ticketPrice(300, true, true)
	receipt, total := formatReceipt("Начало", 2, price)
	fmt.Println("Цена билета:", price)
	fmt.Println(receipt)
	fmt.Println("Общая стоимость:", total)
	fmt.Println("Просмотр разрешён:", isAllowed(18, 16))

	fmt.Println("\nЗадача 2. Контроль доступа к GoFlix")
	if !verifyAccount(18, "premium", false, 16) {
		fmt.Println("Просмотр запрещён")
	}

	fmt.Println("\nЗадача 3. Расписание кинотеатра GoFlix")
	printSchedule()

	fmt.Println("\nЗадача 4. Удаление дубликатов из слайса")
	films := []string{"Дюна", "Матрица", "Дюна", "Начало", "Матрица"}
	fmt.Println("Исходный слайс:", films)
	fmt.Println("Без дубликатов:", removeDuplicates(films))

	var numbers []int
	fmt.Println("Сумма пустого слайса:", sum(numbers))
	ratings := []float64{8.5, 7.2, 9.0, 6.8, 8.1}
	fmt.Println("Рейтинги от 7.5:", filterByRating(ratings, 7.5))
}
